package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetSlidingWindowMeasurementIncreases(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 5},
		{input: util.GetInput(), want: 1275},
	}

	for _, tc := range tests {
		got := GetSlidingWindowMeasurementIncreases(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d increases but got %d", tc.want, got)
		}
	}
}
