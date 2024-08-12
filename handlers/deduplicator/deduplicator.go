package deduplicator

import (
	"ais-stream/models"
	"context"
	"log/slog"
	"sync"
	"time"
)

// Deduplicator stores an indexed hash map of received messages, to identify those
// that are exact duplicates of recently received sentences. The 'timeout' is calculated
// from the time the first message was received - any duplicates within the timeout window
// will be marked accordingly.
//
// Internally, Deduplicator maintains a thread safe hashtable of sync.map[content]<sequence, timestamp>
// where [content] is a fast hash of the nmea string (which by default includes mmsi)
//
// A worker function cleans out the index every 'timeout' interval.Note that sufficient memory
// is required to store 2 x timeout worth of messages.
type Deduplicator struct {
	hashtable  sync.Map
	timeout    time.Duration
	processed  uint64
	duplicates uint64
	discarded  uint64
}

type mapItem struct {
	Timestamp time.Time
	Sequence  uint64
}

func New(timeout time.Duration) *Deduplicator {

	// create a new deduplicator
	d := &Deduplicator{
		hashtable:  sync.Map{},
		timeout:    timeout,
		processed:  0,
		duplicates: 0,
		discarded:  0,
	}
	return d
}

func (d *Deduplicator) IsDuplicate(message models.Message) (isDuplicate bool, id uint64, duplicateId uint64) {

	// get a key for each message - if the key is "" then
	// the message is not of interest and can be discarded
	d.processed++
	key := getKey(message)
	if len(key) == 0 {
		d.discarded++
		return false, d.processed, d.processed
	}

	// test whether the value is present in our sync.map
	value, ok := d.hashtable.Load(key)
	if ok {
		item, ok := value.(*mapItem)
		if ok {
			if item.Timestamp.Add(d.timeout).After(time.Now()) {
				d.duplicates++
				return true, d.processed, item.Sequence
			}
		}
	}

	// if not a duplicate, store the new key-pair with timestamp
	timestamp := time.Unix(message.TagBlock.Time, 0)
	item := mapItem{Timestamp: timestamp, Sequence: d.processed}
	d.hashtable.Store(key, &item)
	return false, d.processed, d.processed
}

// create a worker function to vacuum up stale entries
func (d *Deduplicator) Start(ctx context.Context) {

	// wait for a timer to fire and trigger a vaccuum
	for {
		select {
		case <-ctx.Done():
			slog.Info("deduplicater: stopped worker")
			return
		case <-time.After(d.timeout):

			var (
				counted int       = 0
				deleted int       = 0
				start   time.Time = time.Now()
			)

			d.hashtable.Range(func(key, value interface{}) bool {
				counted++
				item, ok := value.(*mapItem)
				if ok {
					if item.Timestamp.Add(d.timeout).Before(time.Now()) {
						d.hashtable.Delete(key)
						deleted++
					}
				}
				return true
			})
			slog.Info("deduplicater: vacuum",
				"before", counted,
				"deleted", deleted,
				"after", counted-deleted,
				"millis", time.Since(start).Milliseconds(),
				"msg processed", d.processed,
				"msg duplicates", d.duplicates,
			)
		}
	}
}
