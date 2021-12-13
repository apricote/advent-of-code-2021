package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestCountDotsAfterOneFold(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 17},
		{input: util.GetInput(), want: 847},
	}

	for _, tc := range tests {
		got := CountDotsAfterOneFold(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
