/*
Package rtxt converts rich text to HTML.

It supports the following inline formatting:

	**bold**       → <strong>bold</strong>
	//italic//     → <em>italic</em>
	__underline__  → <u>underline</u>
	||highlight||  → <mark>highlight</mark>

Paragraphs are separated by blank lines (double newline).
Single newlines within a paragraph become <br> tags.
HTML special characters are escaped.

	input := "Hello **world**\n\nSecond paragraph"
	output := rtxt.ToHTML(input)
	// "<p>Hello <strong>world</strong></p><p>Second paragraph</p>"
*/
package rtxt
