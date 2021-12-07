package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetDivePositionWithAim(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 900},
		{input: util.GetInput(), want: 1942068080},
	}

	for _, tc := range tests {
		got := GetDivePositionWithAim(tc.input)

		if tc.want != got {
			t.Errorf("Expected dive position %d but got %d", tc.want, got)
		}
	}
}
