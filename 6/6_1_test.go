package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestCountLanternfish(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 5934},
		{input: util.GetInput(), want: 352151},
	}

	for _, tc := range tests {
		got := CountLanternfish(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of lanternfishes to be %d but got %d", tc.want, got)
		}
	}
}
