package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetFinalBingoScoreLosingBoard(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 1924},
		{input: util.GetInput(), want: 23670},
	}

	for _, tc := range tests {
		got := GetFinalBingoScoreLosingBoard(tc.input)

		if tc.want != got {
			t.Errorf("Expected final score for losing board %d but got %d", tc.want, got)
		}
	}
}
