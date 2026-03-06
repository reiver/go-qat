package qat

import (
	"io"
)

// Decoder reads and decodes QAT-formatted question-answer pairs from an input stream.
type Decoder struct {
	scanner *lineScanner
	pending *block // one-block look-ahead buffer
	done    bool
}

// NewDecoder creates a new Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		scanner: newLineScanner(r),
	}
}

// nextBlock returns the next block, using the pending buffer if available.
func (d *Decoder) nextBlock() *block {
	if d.pending != nil {
		b := d.pending
		d.pending = nil
		return b
	}
	return readBlock(d.scanner)
}

// Decode reads the next QA pair from the input.
// Returns io.EOF when no more pairs are available.
func (d *Decoder) Decode(dst *QA) error {
	if d.done {
		return io.EOF
	}

	var question *block

	for {
		b := d.nextBlock()

		if b == nil {
			// EOF
			if question != nil {
				// Q at EOF → emit with empty answer
				dst.Question = question.content
				dst.Answer = ""
				d.done = true
				return nil
			}
			return io.EOF
		}

		switch b.marker {
		case "Q":
			if question != nil {
				// Q-Q: save new Q in pending, emit previous Q with empty answer
				d.pending = b
				dst.Question = question.content
				dst.Answer = ""
				return nil
			}
			question = b

		case "A":
			if question != nil {
				// Q-A: normal pair
				dst.Question = question.content
				dst.Answer = b.content
			} else {
				// A without Q: emit with empty question
				dst.Question = ""
				dst.Answer = b.content
			}
			return nil

		default:
			// Other marker (Note:, Tag:, etc.): skip
			continue
		}
	}
}
