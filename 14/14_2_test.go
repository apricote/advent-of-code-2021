package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCommonElementCountAfter40Rounds(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 2188189693529},
		{input: util.GetInput(), want: 2360298895777},
	}

	for _, tc := range tests {
		got := GetCommonElementCountAfter40Rounds(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
