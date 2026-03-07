package jsonld

import (
	"html"
	"regexp"
	"strings"
)

var (
	paragraphSplitter = regexp.MustCompile(`\n[ \t\n]*\n`)

	boldRegexp      = regexp.MustCompile(`\*\*(.+?)\*\*`)
	italicRegexp    = regexp.MustCompile(`//(.+?)//`)
	underlineRegexp = regexp.MustCompile(`__(.+?)__`)
	highlightRegexp = regexp.MustCompile(`\|\|(.+?)\|\|`)
)

func textToHTML(s string) string {
	if "" == s {
		return ""
	}

	paragraphs := paragraphSplitter.Split(s, -1)

	var buf strings.Builder
	for _, p := range paragraphs {
		p = html.EscapeString(p)

		p = boldRegexp.ReplaceAllString(p, "<strong>$1</strong>")
		p = italicRegexp.ReplaceAllString(p, "<em>$1</em>")
		p = underlineRegexp.ReplaceAllString(p, "<u>$1</u>")
		p = highlightRegexp.ReplaceAllString(p, "<mark>$1</mark>")

		p = strings.ReplaceAll(p, "\n", "<br>")

		buf.WriteString("<p>")
		buf.WriteString(p)
		buf.WriteString("</p>")
	}

	return buf.String()
}
