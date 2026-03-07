package jsonld_test

import (
	"encoding/json"
	"testing"

	"github.com/reiver/go-qat"
	"github.com/reiver/go-qat/jsonld"
)

func TestMarshalFAQPage(t *testing.T) {

	t.Run("single QA pair", func(t *testing.T) {
		qas := []qat.QA{
			{Question: "What is Go?", Answer: "A programming language."},
		}

		data, err := jsonld.MarshalFAQPage(qas)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}

		if got := result["@type"]; got != "FAQPage" {
			t.Errorf("@type = %v, want %q", got, "FAQPage")
		}

		mainEntity, ok := result["mainEntity"].([]interface{})
		if !ok {
			t.Fatalf("mainEntity is not an array: %T", result["mainEntity"])
		}
		if len(mainEntity) != 1 {
			t.Fatalf("mainEntity length = %d, want 1", len(mainEntity))
		}

		q, ok := mainEntity[0].(map[string]interface{})
		if !ok {
			t.Fatalf("mainEntity[0] is not an object")
		}
		if got := q["@type"]; got != "Question" {
			t.Errorf("question @type = %v, want %q", got, "Question")
		}
		if got := q["name"]; got != "What is Go?" {
			t.Errorf("question name = %v, want %q", got, "What is Go?")
		}

		answer, ok := q["acceptedAnswer"].(map[string]interface{})
		if !ok {
			t.Fatalf("acceptedAnswer is not an object")
		}
		if got := answer["@type"]; got != "Answer" {
			t.Errorf("answer @type = %v, want %q", got, "Answer")
		}
		if got := answer["text"]; got != "<p>A programming language.</p>" {
			t.Errorf("answer text = %v, want %q", got, "<p>A programming language.</p>")
		}
	})

	t.Run("multiple QA pairs", func(t *testing.T) {
		qas := []qat.QA{
			{Question: "Q1", Answer: "A1"},
			{Question: "Q2", Answer: "A2"},
			{Question: "Q3", Answer: "A3"},
		}

		data, err := jsonld.MarshalFAQPage(qas)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}

		mainEntity, ok := result["mainEntity"].([]interface{})
		if !ok {
			t.Fatalf("mainEntity is not an array: %T", result["mainEntity"])
		}
		if len(mainEntity) != 3 {
			t.Fatalf("mainEntity length = %d, want 3", len(mainEntity))
		}
	})

	t.Run("answer with formatting produces HTML", func(t *testing.T) {
		qas := []qat.QA{
			{Question: "Formatting?", Answer: "This is **bold** text."},
		}

		data, err := jsonld.MarshalFAQPage(qas)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}

		mainEntity := result["mainEntity"].([]interface{})
		q := mainEntity[0].(map[string]interface{})
		answer := q["acceptedAnswer"].(map[string]interface{})

		expected := "<p>This is <strong>bold</strong> text.</p>"
		if got := answer["text"]; got != expected {
			t.Errorf("answer text = %v, want %q", got, expected)
		}
	})

	t.Run("empty QA slice", func(t *testing.T) {
		data, err := jsonld.MarshalFAQPage([]qat.QA{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}

		if got := result["@type"]; got != "FAQPage" {
			t.Errorf("@type = %v, want %q", got, "FAQPage")
		}
	})

	t.Run("nil QA slice", func(t *testing.T) {
		data, err := jsonld.MarshalFAQPage(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}

		if got := result["@type"]; got != "FAQPage" {
			t.Errorf("@type = %v, want %q", got, "FAQPage")
		}
	})
}
