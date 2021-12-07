package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetFinalBingoScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 4512},
		{input: util.GetInput(), want: 44088},
	}

	for _, tc := range tests {
		got := GetFinalBingoScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected final score %d but got %d", tc.want, got)
		}
	}
}
