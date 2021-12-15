package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestSolveCurrentDayWithTwist(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 315},
		{input: util.GetInput(), want: 2935},
	}

	for _, tc := range tests {
		got := SolveCurrentDayWithTwist(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
