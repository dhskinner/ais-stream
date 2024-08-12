package tcpstream

import (
	"ais-stream/interfaces"
	"ais-stream/sources"
	"context"
	"log/slog"
	"net"
	"sync"
	"time"
)

type tcpClient struct {
	ctx     context.Context
	config  *sources.Config
	handler interfaces.Handler
}

func Client(
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
	c := &tcpClient{ctx: ctx, config: config, handler: h}
	for {

		// run a new worker
		err := c.run()
		if err != nil {
			slog.Error("tcpstream: error", "error", err)
		}

		// on error (and if not cancelled), automatically restart
		select {
		case <-ctx.Done():
			slog.Info("tcpstream: stopped worker")
			return
		default:
			time.Sleep(time.Duration(retry) * time.Second)
		}
	}
}

func (c *tcpClient) run() error {

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
	addr, err := net.ResolveTCPAddr(c.config.Protocol, address)
	if err != nil {
		slog.Info("tcpstream: failed to resolve address", "address", address)
		return err
	}

	// connect!
	conn, err := net.DialTCP(c.config.Protocol, nil, addr)
	if err != nil {
		slog.Info("tcpstream: dial failed", "address", address)
		return err
	}
	slog.Info("tcpstream: connected", "uri", address)

	// defer a function for orderly shutdown
	defer func() {
		conn.Close()
		slog.Info("tcpstream: disconnected", "address", address)
	}()

	slog.Info("tcpstream: starting worker")

worker:
	for {

		// listen for new messages, with a timeout
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		n, err := conn.Read(buffer)

		// check for timeouts
		if err, ok := err.(net.Error); ok && err.Timeout() {
			slog.Info("tcpstream: timeout receiving messages")
			return err
		}

		// check for other errors
		if err != nil {
			slog.Info("tcpstream: error receiving messages")
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

	slog.Info("tcpstream: stopping worker")
	return nil
}
