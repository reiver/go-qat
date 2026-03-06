package qat

import (
	"io"
)

// DecodeAll reads all QA pairs from r and returns them as a slice.
func DecodeAll(r io.Reader) ([]QA, error) {
	decoder := NewDecoder(r)

	var result []QA
	for {
		var qa QA
		err := decoder.Decode(&qa)
		if err == io.EOF {
			return result, nil
		}
		if err != nil {
			return result, err
		}
		result = append(result, qa)
	}
}
