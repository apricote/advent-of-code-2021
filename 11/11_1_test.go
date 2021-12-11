package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestCountFlashesIn100Steps(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 1656},
		{input: util.GetInput(), want: 1608},
	}

	for _, tc := range tests {
		got := CountFlashesIn100Steps(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
