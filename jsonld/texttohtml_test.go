package jsonld

import (
	"testing"
)

func TestTextToHTML(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "empty string",
			input:  "",
			expect: "",
		},
		{
			name:   "plain text",
			input:  "Hello world",
			expect: "<p>Hello world</p>",
		},
		{
			name:   "two paragraphs",
			input:  "First paragraph\n\nSecond paragraph",
			expect: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			name:   "paragraphs separated by whitespace between newlines",
			input:  "First paragraph\n \t\nSecond paragraph",
			expect: "<p>First paragraph</p><p>Second paragraph</p>",
		},
		{
			name:   "bold",
			input:  "**bold**",
			expect: "<p><strong>bold</strong></p>",
		},
		{
			name:   "italic",
			input:  "//italic//",
			expect: "<p><em>italic</em></p>",
		},
		{
			name:   "underline",
			input:  "__underline__",
			expect: "<p><u>underline</u></p>",
		},
		{
			name:   "highlight",
			input:  "||highlight||",
			expect: "<p><mark>highlight</mark></p>",
		},
		{
			name:   "mixed formatting in one paragraph",
			input:  "This is **bold** and //italic// and __underlined__ and ||highlighted||",
			expect: "<p>This is <strong>bold</strong> and <em>italic</em> and <u>underlined</u> and <mark>highlighted</mark></p>",
		},
		{
			name:   "single newline within paragraph becomes br",
			input:  "Line one\nLine two",
			expect: "<p>Line one<br>Line two</p>",
		},
		{
			name:   "HTML special chars escaped",
			input:  "<script>alert('xss')</script>",
			expect: "<p>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</p>",
		},
		{
			name:   "ampersand escaped",
			input:  "AT&T",
			expect: "<p>AT&amp;T</p>",
		},
		{
			name:   "three paragraphs",
			input:  "One\n\nTwo\n\nThree",
			expect: "<p>One</p><p>Two</p><p>Three</p>",
		},
		{
			name:   "formatting across multiple paragraphs",
			input:  "**bold text**\n\n//italic text//",
			expect: "<p><strong>bold text</strong></p><p><em>italic text</em></p>",
		},
		{
			name:   "single newline and double newline mixed",
			input:  "Line one\nLine two\n\nNew paragraph",
			expect: "<p>Line one<br>Line two</p><p>New paragraph</p>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := textToHTML(tt.input)
			if got != tt.expect {
				t.Errorf("textToHTML(%q)\n got: %q\nwant: %q", tt.input, got, tt.expect)
			}
		})
	}
}
