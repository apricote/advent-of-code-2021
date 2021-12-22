package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestCountOnCubes(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetInputFromPath("input_example1.txt"), want: 39},
		{input: util.GetInputFromPath("input_example2.txt"), want: 590784},
		//{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := CountOnCubes(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
