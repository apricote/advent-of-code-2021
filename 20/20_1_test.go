package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestSolveCurrentDay(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 0},
		//{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := SolveCurrentDay(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
