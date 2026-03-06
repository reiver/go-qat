package qat_test

import (
	"strings"
	"testing"

	"github.com/reiver/go-qat"
)

func TestDecodeAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []qat.QA
	}{
		{
			name:  "single inline pair",
			input: "Q: What is Go?\n\nA: A programming language.",
			expect: []qat.QA{
				{Question: "What is Go?", Answer: "A programming language."},
			},
		},
		{
			name:  "multiple inline pairs",
			input: "Q: What is Go?\n\nA: A programming language.\n\nQ: Who made Go?\n\nA: Google.",
			expect: []qat.QA{
				{Question: "What is Go?", Answer: "A programming language."},
				{Question: "Who made Go?", Answer: "Google."},
			},
		},
		{
			name:  "multi-line answer",
			input: "Q: What is a goroutine?\n\nA:\n\tA lightweight thread.\n\t\n\tManaged by the Go runtime.",
			expect: []qat.QA{
				{Question: "What is a goroutine?", Answer: "A lightweight thread.\n\nManaged by the Go runtime."},
			},
		},
		{
			name:  "multi-line question",
			input: "Q:\n\tWhat is:\n\t\n\t5 + 2\n\t\n\t?\n\nA: 7",
			expect: []qat.QA{
				{Question: "What is:\n\n5 + 2\n\n?", Answer: "7"},
			},
		},
		{
			name:  "both multi-line",
			input: "Q:\n\tLine one\n\tLine two\n\nA:\n\tAnswer one\n\tAnswer two",
			expect: []qat.QA{
				{Question: "Line one\nLine two", Answer: "Answer one\nAnswer two"},
			},
		},
		{
			name:  "Q-Q: first Q gets empty answer",
			input: "Q: First\n\nQ: Second\n\nA: Answer for second",
			expect: []qat.QA{
				{Question: "First", Answer: ""},
				{Question: "Second", Answer: "Answer for second"},
			},
		},
		{
			name:  "A without preceding Q",
			input: "A: Orphan answer",
			expect: []qat.QA{
				{Question: "", Answer: "Orphan answer"},
			},
		},
		{
			name:  "Q at EOF gets empty answer",
			input: "Q: Lonely question",
			expect: []qat.QA{
				{Question: "Lonely question", Answer: ""},
			},
		},
		{
			name:  "other markers skipped",
			input: "Note: This is a note\n\nQ: Real question\n\nTag: sometag\n\nA: Real answer",
			expect: []qat.QA{
				{Question: "Real question", Answer: "Real answer"},
			},
		},
		{
			name:  "unicode markers skipped",
			input: "Q: Hello\n\nA: World",
			expect: []qat.QA{
				{Question: "Hello", Answer: "World"},
			},
		},
		{
			name:  "marker with hyphen",
			input: "Meta-Info: some metadata\n\nQ: Question\n\nA: Answer",
			expect: []qat.QA{
				{Question: "Question", Answer: "Answer"},
			},
		},
		{
			name:  "marker with underscore and digits",
			input: "Tag_2: something\n\nQ: Question\n\nA: Answer",
			expect: []qat.QA{
				{Question: "Question", Answer: "Answer"},
			},
		},
		{
			name:   "empty input",
			input:  "",
			expect: nil,
		},
		{
			name:   "only whitespace",
			input:  "   \n\n   \n",
			expect: nil,
		},
		{
			name:  "trailing blank lines in multi-line content trimmed",
			input: "Q: Question\n\nA:\n\tLine one\n\t\n\t\n",
			expect: []qat.QA{
				{Question: "Question", Answer: "Line one"},
			},
		},
		{
			name:  "deeper indentation preserved",
			input: "Q: Example\n\nA:\n\tCode example:\n\t\n\t\tgo fn()\n\t\n\tEnd.",
			expect: []qat.QA{
				{Question: "Example", Answer: "Code example:\n\n\tgo fn()\n\nEnd."},
			},
		},
		{
			name:  "whitespace-only lines between blocks ignored",
			input: "Q: First\n   \n\nA: Answer\n   \n\nQ: Second\n\nA: Answer2",
			expect: []qat.QA{
				{Question: "First", Answer: "Answer"},
				{Question: "Second", Answer: "Answer2"},
			},
		},
		{
			name:  "blank lines between multi-line content preserved",
			input: "A:\n\tParagraph one.\n\t\n\t\n\tParagraph two.",
			expect: []qat.QA{
				{Question: "", Answer: "Paragraph one.\n\n\nParagraph two."},
			},
		},
		{
			name:  "readme example 1: inline pairs",
			input: "Q: What is a goroutine in Go?\n\nA: A goroutine is a lightweight thread managed by the Go runtime.\n\nQ: What is the main package?\n\nA: The main package defines an executable program. The program starts execution in main().",
			expect: []qat.QA{
				{Question: "What is a goroutine in Go?", Answer: "A goroutine is a lightweight thread managed by the Go runtime."},
				{Question: "What is the main package?", Answer: "The main package defines an executable program. The program starts execution in main()."},
			},
		},
		{
			name:  "readme example 2: multi-line answer",
			input: "Q: What is a goroutine in Go?\n\nA:\n\tA goroutine is a lightweight thread managed by the Go runtime.\n\t\n\tHere is an example:\n\t\n\t\tgo fn()\n\nQ: What is the main package?\n\nA: The main package defines an executable program. The program starts execution in main().",
			expect: []qat.QA{
				{Question: "What is a goroutine in Go?", Answer: "A goroutine is a lightweight thread managed by the Go runtime.\n\nHere is an example:\n\n\tgo fn()"},
				{Question: "What is the main package?", Answer: "The main package defines an executable program. The program starts execution in main()."},
			},
		},
		{
			name:  "readme example 3: multi-line questions",
			input: "Q:\n\tWhat is:\n\t\n\t5 + 2\n\t\n\t?\n\nA: 7\n\nQ:\n\tWhat is:\n\t\n\t5 - 2\n\t\n\t?\n\nA: 3",
			expect: []qat.QA{
				{Question: "What is:\n\n5 + 2\n\n?", Answer: "7"},
				{Question: "What is:\n\n5 - 2\n\n?", Answer: "3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := qat.DecodeAll(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(result) != len(tt.expect) {
				t.Fatalf("got %d pairs, want %d\ngot:  %#v\nwant: %#v", len(result), len(tt.expect), result, tt.expect)
			}

			for i := range tt.expect {
				if result[i].Question != tt.expect[i].Question {
					t.Errorf("pair %d: question mismatch\n got: %q\nwant: %q", i, result[i].Question, tt.expect[i].Question)
				}
				if result[i].Answer != tt.expect[i].Answer {
					t.Errorf("pair %d: answer mismatch\n got: %q\nwant: %q", i, result[i].Answer, tt.expect[i].Answer)
				}
			}
		})
	}
}
