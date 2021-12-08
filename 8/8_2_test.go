package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetOutputSum(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 61229},
		{input: util.GetInput(), want: 1007675},
	}

	for _, tc := range tests {
		got := GetOutputSum(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
