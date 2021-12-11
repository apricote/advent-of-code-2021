package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetFirstDayWithSynchronizedFlashes(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 195},
		{input: util.GetInput(), want: 214},
	}

	for _, tc := range tests {
		got := GetFirstDayWithSynchronizedFlashes(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
