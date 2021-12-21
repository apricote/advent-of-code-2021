package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestFindAllPossibleVelocities(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 112},
		{input: util.GetInput(), want: 1919},
	}

	for _, tc := range tests {
		got := FindAllPossibleVelocities(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
