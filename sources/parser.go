package sources

import (
	"ais-stream/interfaces"
	"ais-stream/models"
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
)

const (
	PARSER_STATS_INTERVAL time.Duration = 1 * time.Minute
)

type State int

const (
	idle State = iota
	tagblock
	waitingForSentence
	sentence
)

// The sentence parser separates bytes into sentences, then decodes the AIS message
// when a full sentence is received. For sentences that span across packets, the
// timestamp of the last packet is used for simplicity. Note you should only ever
// assign one parser to one input source (1:1)
type Parser struct {
	name         string
	state        State
	source       chan byte
	buffer       chan byte
	handler      interfaces.Handler
	pending      *models.Sentence
	codec        *aisnmea.NMEACodec
	messageCount uint64
	errorCount   uint64
	lastCount    uint64
	lastStats    time.Time
	verbose      bool
}

func NewParser(hd interfaces.Handler, name string, verbose bool) *Parser {

	p := &Parser{
		name:         name,
		source:       make(chan byte, 4096),
		buffer:       make(chan byte, 4096),
		handler:      hd,
		pending:      &models.Sentence{},
		state:        idle,
		codec:        aisnmea.NMEACodecNew(ais.CodecNew(false, false)),
		messageCount: 0,
		errorCount:   0,
		lastCount:    0,
		lastStats:    time.Now(),
		verbose:      verbose,
	}

	return p
}

// flag to set verbose mode
func (p *Parser) SetVerbose(verbose bool) {
	p.verbose = verbose
}

// print some statistics
func (p *Parser) PrintStats() {
	count := p.messageCount - p.lastCount
	durn := time.Since(p.lastStats)
	freq := fmt.Sprintf("%.1f", float32(count)/float32(durn.Seconds()))
	p.lastCount = p.messageCount
	p.lastStats = time.Now()
	slog.Info("parser", "name", p.name, "messages", p.messageCount, "rate/sec", freq)
}

// Accepts and copies bytes from an input source into a buffer for processing
func (p *Parser) AddBytes(in []byte) {
	for i := 0; i < len(in); i++ {
		p.source <- in[i]
	}
}

// Accepts and copies strings from an input source into a buffer for processing
// Note this automatically adds an ‘\r\n’ delimiter between strings
func (p *Parser) AddStrings(in []string) {
	for i := 0; i < len(in); i++ {
		p.AddBytes([]byte(in[i]))
		p.source <- '\r'
		p.source <- '\n'
	}
}

// Processes incoming bytes to separate discrete sentences
func (p *Parser) Process(ctx context.Context) {

	p.reset()

	ticker := time.NewTicker(PARSER_STATS_INTERVAL)
	defer ticker.Stop()

worker:
	for {

		select {

		// check for cancel signal
		case <-ctx.Done():
			break worker

		// print some stats
		case <-ticker.C:
			p.PrintStats()

		// receive incoming bytes
		case c := <-p.source:

			// ignore null bytes
			if c == 0x00 {
				continue
			}

			// is each byte within limits between 0x20 (space) to 0x7e (~)
			if c > 0x7E ||
				c < 0x0A ||
				(c > 0x0A && c < 0x0D) ||
				(c > 0x0D && c < 0x20) {
				p.reset()
			}

			// process each byte via a simple state machine
			switch p.state {

			case tagblock:

				switch c {
				// have we found the tagblock end delimiter <\>
				case '\\':
					p.buffer <- c
					p.pending.TagBlock = p.bufferAsString()
					p.state = waitingForSentence

				// have we found something we shouldn't have
				case '$', '!', '\n', '\r':
					p.reset()

				// otherwise buffer the tagblock bytes
				default:
					p.buffer <- c
				}

			case waitingForSentence:

				switch c {
				// the first character must be an nmea start delimiter
				case '!', '$':
					p.buffer <- c
					p.state = sentence

				// not found, so discard and reset
				default:
					p.reset()
				}

			case sentence:

				switch c {

				// output a complete sentence when an end delimiter <\n> is found
				case '\n':
					p.buffer <- c
					p.pending.Content = p.bufferAsString()
					p.addTagBlock()
					p.decode(p.pending)
					p.reset()

				// have we found something we shouldn't have
				case '$', '!', '\\':
					p.reset()

				// otherwise buffer the sentence content
				default:
					p.buffer <- c
				}

			default:

				switch c {

				// look for the start of a new tagblock
				case '\\':
					p.state = tagblock
					p.buffer <- c

				// look for the start of a new sentence
				case '!', '$':
					p.state = sentence
					p.buffer <- c

				// discard and move on
				default:

				}
			}
		}
	}
}

// output the current buffer as a string
func (p *Parser) bufferAsString() string {

	buf := make([]byte, len(p.buffer))
	for len(p.buffer) > 0 {
		b := <-p.buffer
		buf = append(buf, b)
	}
	return string(bytes.Trim(buf, "\x00"))

}

// clear the buffer
func (p *Parser) reset() {

	for len(p.buffer) > 0 {
		<-p.buffer
	}
	p.pending = &models.Sentence{}
	p.state = idle

}

// add optional tags
func (p *Parser) addTagBlock() {

	// is there an existing tagblock? see if it has anything useful
	tagblock, err := parseTagBlock(p.pending.TagBlock)
	if err != nil {
		slog.Debug("error parsing tagblock", "input", p.pending.TagBlock, "error", err)
	}

	// add in missing time and source fields
	if tagblock.Time == 0 {
		tagblock.Time = time.Now().Unix()
	}
	if len(tagblock.Source) == 0 {
		tagblock.Source = p.name
	}

	// add a 2 character source identifier into the text field
	tagblock.Text = fmt.Sprintf("<src %s>%s", p.name[0:2], tagblock.Text)

	p.pending.TagBlock = tagBlockAsString(&tagblock)
}

func (p *Parser) decode(sentence *models.Sentence) error {

	p.messageCount++

	reassembled := sentence.AsString()
	if len(reassembled) == 0 {
		return nil
	}

	message, err := p.codec.ParseSentence(reassembled)
	if err != nil {
		p.errorCount++
		slog.Debug("error decoding sentence",
			"input", sentence.AsString(),
			"error", err,
			"count", fmt.Sprintf("%d of %d", p.errorCount, p.messageCount))
	} else if message != nil {

		if p.verbose {
			slog.Info("parser", "name", p.name, "tags", message.TagBlock)
		}
		p.handler.Message(models.Message(message))

	}
	return err
}
