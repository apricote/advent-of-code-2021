package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestCountLanternfish256(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 26984457539},
		{input: util.GetInput(), want: 1601616884019},
	}

	for _, tc := range tests {
		got := CountLanternfish256(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of lanternfishes after 256 days to be %d but got %d", tc.want, got)
		}
	}
}
