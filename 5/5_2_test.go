package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetOverlappingVentsWithDiagonals(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 12},
		{input: util.GetInput(), want: 22335},
	}

	for _, tc := range tests {
		got := GetOverlappingVentsWithDiagonals(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of overlapping vents (with diagonals) to be %d but got %d", tc.want, got)
		}
	}
}
