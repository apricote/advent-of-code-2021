package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestFindTimeToLand(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 58},
		{input: util.GetInput(), want: 453},
	}

	for _, tc := range tests {
		got := FindTimeToLand(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
