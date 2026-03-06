package qat_test

import (
	"io"
	"strings"
	"testing"

	"github.com/reiver/go-qat"
)

func TestDecoder_Decode(t *testing.T) {
	t.Run("streaming decode returns pairs one at a time", func(t *testing.T) {
		input := "Q: First\n\nA: Answer1\n\nQ: Second\n\nA: Answer2\n"

		dec := qat.NewDecoder(strings.NewReader(input))

		var qa qat.QA

		// First pair
		err := dec.Decode(&qa)
		if err != nil {
			t.Fatalf("first Decode: unexpected error: %v", err)
		}
		if qa.Question != "First" || qa.Answer != "Answer1" {
			t.Errorf("first pair: got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "First", "Answer1")
		}

		// Second pair
		err = dec.Decode(&qa)
		if err != nil {
			t.Fatalf("second Decode: unexpected error: %v", err)
		}
		if qa.Question != "Second" || qa.Answer != "Answer2" {
			t.Errorf("second pair: got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "Second", "Answer2")
		}

		// EOF
		err = dec.Decode(&qa)
		if err != io.EOF {
			t.Errorf("expected io.EOF, got %v", err)
		}
	})

	t.Run("repeated EOF calls return EOF", func(t *testing.T) {
		dec := qat.NewDecoder(strings.NewReader(""))

		var qa qat.QA
		for i := 0; i < 3; i++ {
			err := dec.Decode(&qa)
			if err != io.EOF {
				t.Errorf("call %d: expected io.EOF, got %v", i, err)
			}
		}
	})

	t.Run("Q at EOF emits then EOF on next call", func(t *testing.T) {
		dec := qat.NewDecoder(strings.NewReader("Q: Lonely"))

		var qa qat.QA
		err := dec.Decode(&qa)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if qa.Question != "Lonely" || qa.Answer != "" {
			t.Errorf("got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "Lonely", "")
		}

		err = dec.Decode(&qa)
		if err != io.EOF {
			t.Errorf("expected io.EOF, got %v", err)
		}
	})

	t.Run("Q-Q-A produces two pairs", func(t *testing.T) {
		input := "Q: First\n\nQ: Second\n\nA: Answer"
		dec := qat.NewDecoder(strings.NewReader(input))

		var qa qat.QA

		err := dec.Decode(&qa)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if qa.Question != "First" || qa.Answer != "" {
			t.Errorf("first pair: got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "First", "")
		}

		err = dec.Decode(&qa)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if qa.Question != "Second" || qa.Answer != "Answer" {
			t.Errorf("second pair: got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "Second", "Answer")
		}

		err = dec.Decode(&qa)
		if err != io.EOF {
			t.Errorf("expected io.EOF, got %v", err)
		}
	})

	t.Run("interleaved other markers", func(t *testing.T) {
		input := "Note: skip this\n\nQ: Real\n\nTag: also skip\n\nA: Answer\n\nNote: trailing"
		dec := qat.NewDecoder(strings.NewReader(input))

		var qa qat.QA
		err := dec.Decode(&qa)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if qa.Question != "Real" || qa.Answer != "Answer" {
			t.Errorf("got {%q, %q}, want {%q, %q}", qa.Question, qa.Answer, "Real", "Answer")
		}

		err = dec.Decode(&qa)
		if err != io.EOF {
			t.Errorf("expected io.EOF, got %v", err)
		}
	})
}
