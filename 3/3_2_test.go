package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetLifeSupportRating(t *testing.T) {
	type test struct {
		input string
		want  int64
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 230},
		{input: util.GetInput(), want: 7041258},
	}

	for _, tc := range tests {
		got := GetLifeSupportRating(tc.input)

		if tc.want != got {
			t.Errorf("Expected life support rating %d but got %d", tc.want, got)
		}
	}
}
