package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCompletionScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 288957},
		{input: util.GetInput(), want: 2360030859},
	}

	for _, tc := range tests {
		got := GetCompletionScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
