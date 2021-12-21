package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetDiracQuantumDiceGameScore(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 444356092776315},
		{input: util.GetInput(), want: 133029050096658},
	}

	for _, tc := range tests {
		got := GetDiracQuantumDiceGameScore(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
