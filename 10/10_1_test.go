package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetSyntaxErrorScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 26397},
		{input: util.GetInput(), want: 321237},
	}

	for _, tc := range tests {
		got := GetSyntaxErrorScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
