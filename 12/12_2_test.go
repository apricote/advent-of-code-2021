package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCaveTimesWithAdditionalTime(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetInputFromPath("input_example1.txt"), want: 36},
		{input: util.GetInputFromPath("input_example2.txt"), want: 103},
		{input: util.GetInputFromPath("input_example3.txt"), want: 3509},
		{input: util.GetInput(), want: 85062},
	}

	for _, tc := range tests {
		got := GetCaveTimesWithAdditionalTime(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
