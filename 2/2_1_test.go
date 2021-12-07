package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetDivePosition(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 150},
		{input: util.GetInput(), want: 2039912},
	}

	for _, tc := range tests {
		got := GetDivePosition(tc.input)

		if tc.want != got {
			t.Errorf("Expected dive position %d but got %d", tc.want, got)
		}
	}
}
