package qat

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// isMarkerLine checks whether a line is a marker line.
// A marker line starts at column 0 with a Unicode letter, followed by zero or more
// Unicode letters, digits, underscores, or hyphens, then a colon.
// Returns the marker name and true if it matches, or ("", false) otherwise.
func isMarkerLine(line string) (string, bool) {
	if len(line) == 0 {
		return "", false
	}

	r, size := utf8.DecodeRuneInString(line)
	if r == utf8.RuneError || !unicode.IsLetter(r) {
		return "", false
	}

	i := size
	for i < len(line) {
		r, size = utf8.DecodeRuneInString(line[i:])
		if r == ':' {
			return line[:i], true
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			i += size
			continue
		}
		return "", false
	}
	return "", false
}

// readBlock reads the next marker+content block from the scanner.
// Returns nil when no more blocks are available (EOF).
func readBlock(ls *lineScanner) *block {
	// Step 1: skip non-marker lines until we find a marker line or hit EOF
	var marker string
	var remainder string
	for {
		line, ok := ls.readLine()
		if !ok {
			return nil
		}
		m, isMarker := isMarkerLine(line)
		if isMarker {
			marker = m
			// remainder is everything after "Marker:"
			remainder = line[len(marker)+1:] // +1 for the colon
			break
		}
	}

	// Step 3: inline vs multi-line
	trimmed := strings.TrimSpace(remainder)
	if trimmed != "" {
		return &block{
			marker:  marker,
			content: trimmed,
		}
	}

	// Step 4: multi-line content collection
	var lines []string
	for {
		line, ok := ls.peek()
		if !ok {
			break
		}

		// Empty or whitespace-only line: preserve as blank in content
		if strings.TrimSpace(line) == "" {
			ls.readLine() // consume
			lines = append(lines, "")
			continue
		}

		// Line starting with tab: strip one leading tab
		if line[0] == '\t' {
			ls.readLine() // consume
			lines = append(lines, line[1:])
			continue
		}

		// Non-whitespace at column 0: next marker or non-marker — stop
		break
	}

	// Step 5: trim trailing blank lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return &block{
		marker:  marker,
		content: strings.Join(lines, "\n"),
	}
}
