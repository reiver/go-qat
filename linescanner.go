package qat

import (
	"bufio"
	"io"
)

// lineScanner wraps a bufio.Scanner to provide one-line peek/unread capability.
type lineScanner struct {
	scanner  *bufio.Scanner
	peeked   bool
	peekLine string
	peekOK   bool
}

func newLineScanner(r io.Reader) *lineScanner {
	return &lineScanner{
		scanner: bufio.NewScanner(r),
	}
}

// readLine returns the next line and true, or ("", false) at EOF.
func (ls *lineScanner) readLine() (string, bool) {
	if ls.peeked {
		ls.peeked = false
		return ls.peekLine, ls.peekOK
	}
	if ls.scanner.Scan() {
		return ls.scanner.Text(), true
	}
	return "", false
}

// peekLine returns the next line without consuming it.
func (ls *lineScanner) peek() (string, bool) {
	if ls.peeked {
		return ls.peekLine, ls.peekOK
	}
	if ls.scanner.Scan() {
		ls.peeked = true
		ls.peekLine = ls.scanner.Text()
		ls.peekOK = true
		return ls.peekLine, ls.peekOK
	}
	ls.peeked = true
	ls.peekLine = ""
	ls.peekOK = false
	return ls.peekLine, ls.peekOK
}
