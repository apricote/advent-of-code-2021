package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetOverlappingVents(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 5},
		{input: util.GetInput(), want: 6397},
	}

	for _, tc := range tests {
		got := GetOverlappingVents(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of overlapping vents to be %d but got %d", tc.want, got)
		}
	}
}
