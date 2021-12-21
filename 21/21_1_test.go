package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetDiracDiceGameScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 739785},
		{input: util.GetInput(), want: 805932},
	}

	for _, tc := range tests {
		got := GetDiracDiceGameScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
