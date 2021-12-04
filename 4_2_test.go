package main

import "testing"

func TestGetFinalBingoScoreLosingBoard(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay4ExampleInput(), want: 1924},
		{input: GetDay4Input(), want: 23670},
	}

	for _, tc := range tests {
		got := GetFinalBingoScoreLosingBoard(tc.input)

		if tc.want != got {
			t.Errorf("Expected final score for losing board %d but got %d", tc.want, got)
		}
	}
}
