package main

import "testing"

func TestGetFinalBingoScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay4ExampleInput(), want: 4512},
		{input: GetDay4Input(), want: 44088},
	}

	for _, tc := range tests {
		got := GetFinalBingoScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected final score %d but got %d", tc.want, got)
		}
	}
}
