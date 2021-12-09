package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetBasins(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 1134},
		{input: util.GetInput(), want: 847044},
	}

	for _, tc := range tests {
		got := GetBasins(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
