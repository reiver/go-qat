package qat

// block represents a single marker-content unit parsed from QAT text.
// For example, "Q: hello" produces block{marker: "Q", content: "hello"}.
type block struct {
	marker  string
	content string
}
