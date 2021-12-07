package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetMinimumFuelCost(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 37},
		{input: util.GetInput(), want: 357353},
	}

	for _, tc := range tests {
		got := GetMinimumFuelCost(tc.input)

		if tc.want != got {
			t.Errorf("Expected fuel cost of %d but got %d", tc.want, got)
		}
	}
}
