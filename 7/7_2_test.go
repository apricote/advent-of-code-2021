package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetMinimumFuelCostWithRaisingPrices(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 168},
		{input: util.GetInput(), want: 104822130},
	}

	for _, tc := range tests {
		got := GetMinimumFuelCostWithRaisingPrices(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
