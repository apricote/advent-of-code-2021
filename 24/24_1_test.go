package main

import (
	"testing"
)

func TestMONAD(t *testing.T) {
	type test struct {
		input int
		want  bool
	}

	tests := []test{
		{input: 13579246899999, want: true},
	}

	for _, tc := range tests {
		got := MONAD(tc.input)

		if tc.want != got {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}

func TestBruteForceModelNumber(t *testing.T) {
	modelNumber := BruteForceModelNumber()
	if modelNumber != 0 {
		t.Errorf("Expected 0 but got %v", modelNumber)
	}

}
