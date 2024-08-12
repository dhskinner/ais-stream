package defaulthandler

import (
	"ais-stream/handlers"
	"ais-stream/handlers/deduplicator"
	"ais-stream/interfaces"
	"ais-stream/models"
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

// Default handler for incoming sentences
type DefaultHandler struct {
	queue            chan models.Message
	duplicateTimeout time.Duration
	writeToFile      bool
}

func New(duplicateTimeout time.Duration, writeToFile bool) *DefaultHandler {

	p := &DefaultHandler{
		queue:            make(chan models.Message, 100),
		duplicateTimeout: duplicateTimeout,
		writeToFile:      writeToFile,
	}
	return p

}

func (p *DefaultHandler) Message(message models.Message) error {

	if message == nil {
		return fmt.Errorf("error cannot add nil message")
	}
	p.queue <- message
	return nil

}

// Processes incoming sentences - default handler just outputs these to the console
func (p *DefaultHandler) Process(ctx context.Context, wg *sync.WaitGroup) {

	// tell the caller we've stopped
	defer wg.Done()

	// create a deduplicator to tag repeated messages
	dedup := deduplicator.New(p.duplicateTimeout)
	go dedup.Start(ctx)

	// create a new file to write out log entries
	var writer *bufio.Writer = nil
	if p.writeToFile {
		f, err := os.Create("log.txt")
		if err != nil {
			return
		}
		writer = bufio.NewWriter(f)
		defer func() {
			if writer != nil {
				writer.Flush()
			}
			f.Close()
		}()
	}

worker:
	for {
		select {

		// check for cancel signal
		case <-ctx.Done():
			break worker

		// receive incoming sentences
		case message := <-p.queue:

			// if the message is not of interest, ignore it and move on
			isDuplicate, id, _ := dedup.IsDuplicate(message)

			if isDuplicate {
				handlers.Print(id, handlers.Duplicate, message)
			} else {
				handlers.Print(id, handlers.OK, message)
			}

			// if writer != nil {
			// 	_, err := writer.WriteString(out)
			// 	if err != nil {
			// 		return
			// 	}
			// }
		}
	}
}

func (p *DefaultHandler) Record(mmsi uint32) (*interfaces.Record, error) {
	return nil, fmt.Errorf("not implemented")
}
