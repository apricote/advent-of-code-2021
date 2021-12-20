package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCommonElementCount(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 1588},
		{input: util.GetInput(), want: 2194},
	}

	for _, tc := range tests {
		got := GetCommonElementCount(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
