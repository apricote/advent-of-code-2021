package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetMeasurementIncreases(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 7},
		{input: util.GetInput(), want: 1233},
	}

	for _, tc := range tests {
		got := GetMeasurementIncreases(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d increases but got %d", tc.want, got)
		}
	}
}
