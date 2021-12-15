package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetCostToAvoidChitons(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 40},
		{input: util.GetInput(), want: 602},
	}

	for _, tc := range tests {
		got := GetCostToAvoidChitons(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
