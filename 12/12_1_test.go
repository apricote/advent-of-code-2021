package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCavePaths(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetInputFromPath("input_example1.txt"), want: 10},
		{input: util.GetInputFromPath("input_example2.txt"), want: 19},
		{input: util.GetInputFromPath("input_example3.txt"), want: 226},
		{input: util.GetInput(), want: 3485},
	}

	for _, tc := range tests {
		got := GetCavePaths(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
