package udpstream

import (
	"ais-stream/interfaces"
	"ais-stream/sources"
	"context"
	"log/slog"
	"net"
	"sync"
	"time"
)

type udpServer struct {
	ctx     context.Context
	config  *sources.Config
	handler interfaces.Handler
}

func Server(
	ctx context.Context,
	wg *sync.WaitGroup,
	config *sources.Config,
	h interfaces.Handler,
) {

	// tell the caller we've stopped
	defer wg.Done()

	// the following must be present
	retry, err := config.GetRetry()
	if err != nil {
		return
	}

	// service the message stream (blocking)
	c := &udpServer{ctx: ctx, config: config, handler: h}
	for {

		// run a new worker
		err := c.run()
		if err != nil {
			slog.Error("udpstream: error", "error", err)
		}

		// on error (and if not cancelled), automatically restart
		select {
		case <-ctx.Done():
			slog.Info("udpstream: stopped worker")
			return
		default:
			time.Sleep(time.Duration(retry) * time.Second)
		}
	}
}

func (c *udpServer) run() error {

	// get some environment variables - the following keys must be present:
	address, err := c.config.GetAddress()
	if err != nil {
		return err
	}

	timeout, err := c.config.GetTimeout()
	if err != nil {
		return err
	}

	// create a buffer for incoming packet data
	buffer := make([]byte, 4096)

	// create a parser to extract complete sentences
	parser := sources.NewParser(c.handler, c.config.Name)
	go parser.Process(c.ctx)

	// check the given uri/address is valid
	addr, err := net.ResolveUDPAddr(c.config.Protocol, address)
	if err != nil {
		slog.Info("udpstream: failed to resolve address", "address", address)
		return err
	}

	// connect!
	conn, err := net.ListenUDP(c.config.Protocol, addr)
	if err != nil {
		slog.Info("udpstream: dial failed", "address", address)
		return err
	}
	slog.Info("udpstream: connected", "address", address)

	// defer a function for orderly shutdown
	defer func() {
		conn.Close()
		slog.Info("udpstream: disconnected", "address", address)
	}()

	slog.Info("udpstream: starting worker")

worker:
	for {

		// listen for new messages, with a timeout
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		n, _, err := conn.ReadFromUDP(buffer)

		// check for timeouts
		if err, ok := err.(net.Error); ok && err.Timeout() {
			slog.Info("udpstream: timeout receiving messages")
			return err
		}

		// check for other errors
		if err != nil {
			slog.Info("udpstream: error receiving messages")
			return err
		}

		// parse packets into sentences
		if n > 0 {
			parser.AddBytes(buffer)
		}

		// check for cancel signal
		select {
		case <-c.ctx.Done():
			break worker
		default:
		}
	}

	slog.Info("udpstream: stopping worker")
	return nil
}
