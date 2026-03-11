/*
Package qat decodes the QAT (Question-Answer Text) format.

The QAT format uses marker lines to delimit question-answer pairs:

	Q: What is Go?
	A: A programming language.

Content can also span multiple lines using tab indentation:

	Q:
		What is Go?
	A:
		Go is a statically typed,
		compiled programming language
		designed at Google.

Markers are lines starting with a letter followed by a colon (e.g. "Q:", "A:", "TITLE:", "HEADLINE:").
Only Q and A markers produce QA pairs; all others are skipped.

Use [NewDecoder] for streaming access or [DecodeAll] to read everything at once:

	decoder := qat.NewDecoder(reader)
	var qa qat.QA
	for decoder.Decode(&qa) == nil {
		fmt.Println(qa.Question, qa.Answer)
	}
*/
package qat
