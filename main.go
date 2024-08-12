package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"time"

	"ais-stream/atlas"
	"ais-stream/handlers/mongohandler"
	"ais-stream/interfaces"
	"ais-stream/sources"
	"ais-stream/sources/aisstream"
	"ais-stream/sources/tcpstream"
	"ais-stream/sources/udpstream"
)

const (
	WORKER_CLOSE_TIMEOUT time.Duration = 5 * time.Second
	RUN_AISSTREAM        bool          = false // <- normally true
	RUN_WALNUT_TCP       bool          = false // <- normally true
	RUN_WALNUT_UDP       bool          = false
	RUN_AMSA             bool          = true // <- normally true
	RUN_AISHUB           bool          = false
	RUN_TICKER           bool          = false
)

func main() {

	// set up default logging
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	slog.SetDefault(slog.New(logHandler))
	slog.Info("main: starting workers")

	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	// add a mongodb handler for received sentences
	wg.Add(1)
	mongoWorker := mongohandler.New(
		30*time.Second,
		"LOCAL_MONGODB_CONNECTION",
	)
	go mongoWorker.Process(ctx, &wg, "australia")

	// add a worker to aggregate data and send it to mongodb atlas
	atlasWorker := atlas.New()
	go atlasWorker.Process(ctx, &wg)

	// add a simple ticker
	if RUN_TICKER {
		wg.Add(1)
		go ticker(ctx, &wg, 10*time.Second)
	}

	// add a local ais receiver (tcp)
	if RUN_WALNUT_TCP {
		wg.Add(1)
		go tcpstream.Client(
			ctx,
			&wg,
			&sources.Config{
				Name:           "Bayside",
				Protocol:       "tcp",
				AddressKey:     "AISLOCAL_URI",
				TimeoutSecsKey: "AISLOCAL_TIMEOUT_SECS",
				RetrySecsKey:   "AISLOCAL_RETRY_SECS",
			},
			interfaces.Handler(mongoWorker),
		)
	}

	// add a local ais receiver (udp)
	if RUN_WALNUT_UDP {
		wg.Add(1)
		go udpstream.Server(
			ctx,
			&wg,
			&sources.Config{
				Name:           "Bayside",
				Protocol:       "udp",
				AddressKey:     "UDP_URI",
				TimeoutSecsKey: "UDP_TIMEOUT_SECS",
				RetrySecsKey:   "UDP_RETRY_SECS",
			},
			interfaces.Handler(mongoWorker),
		)
	}

	// add amsa
	if RUN_AMSA {
		wg.Add(1)
		go tcpstream.Client(
			ctx,
			&wg,
			&sources.Config{
				Name:           "Amsa",
				Protocol:       "tcp",
				AddressKey:     "AMSA_URI",
				TimeoutSecsKey: "AMSA_TIMEOUT_SECS",
				RetrySecsKey:   "AMSA_RETRY_SECS",
				Verbose:        true,
			},
			interfaces.Handler(mongoWorker),
		)
	}

	// add aisstream
	if RUN_AISSTREAM {
		wg.Add(1)
		go aisstream.Client(
			ctx,
			&wg,
			&sources.Config{
				Name:           "Aisstream",
				AddressKey:     "AISSTREAM_URI",
				ApiKey:         "AISSTREAM_API_KEY",
				TimeoutSecsKey: "AISSTREAM_TIMEOUT_SECS",
				RetrySecsKey:   "AISSTREAM_RETRY_SECS",
				Boundary: &sources.BoundaryConfig{
					// Lat1Key: "AISSTREAM_BOUNDARY_QLD_LAT1",
					// Lat2Key: "AISSTREAM_BOUNDARY_QLD_LAT2",
					// Lon1Key: "AISSTREAM_BOUNDARY_QLD_LON1",
					// Lon2Key: "AISSTREAM_BOUNDARY_QLD_LON2",
					Lat1Key: "AISSTREAM_BOUNDARY_LAT1",
					Lat2Key: "AISSTREAM_BOUNDARY_LAT2",
					Lon1Key: "AISSTREAM_BOUNDARY_LON1",
					Lon2Key: "AISSTREAM_BOUNDARY_LON2",
				},
			},
			interfaces.Handler(mongoWorker),
		)
	}

	// add aishub
	if RUN_AISHUB {
		wg.Add(1)
		go tcpstream.Client(
			ctx,
			&wg,
			&sources.Config{
				Name:           "Aishub",
				AddressKey:     "AISHUB_URI",
				TimeoutSecsKey: "AISHUB_TIMEOUT_SECS",
				RetrySecsKey:   "AISHUB_RETRY_SECS",
			},
			interfaces.Handler(mongoWorker),
		)
	}

	// listen for <ctl>c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	slog.Info("main: <ctrl>c - shutting down")

	// tell the goroutines to stop
	slog.Info("main: stopping workers")
	cancel()

	// and wait for them to reply back
	if waitTimeout(&wg, WORKER_CLOSE_TIMEOUT) {
		slog.Info("main: timed out waiting for workers to close")
	} else {
		slog.Info("main: all workers were closed gracefully")
	}
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func ticker(ctx context.Context, wg *sync.WaitGroup, interval time.Duration) {

	defer wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case now := <-ticker.C:
			fmt.Printf("ticker: tick %s\n", now.UTC().Format("20060102-150405.000000000"))
		case <-ctx.Done():
			fmt.Println("ticker: stopping")
			return
		}
	}
}
