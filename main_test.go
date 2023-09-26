package main

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "lake", output: "aekl"},
		{input: "alfie", output: "aefil"},
		{input: "lavender", output: "adelnrv"},
		{input: "kale", output: "aekl"},
	}

	for _, tt := range tests {
		h := hash(tt.input)
		if h != tt.output {
			t.Fatalf("hash mismatch: expected=%s got=%s", tt.output, h)
		}
	}
}
