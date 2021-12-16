package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestSolvePacketEquation(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "C200B40A82", want: 3},
		{input: "04005AC33890", want: 54},
		{input: "880086C3E88112", want: 7},
		{input: "CE00C43D881120", want: 9},
		{input: "D8005AC2A8F0", want: 1},
		{input: "F600BC2D8F", want: 0},
		{input: "9C005AC2F8F0", want: 0},
		{input: "9C0141080250320F1802104A08", want: 1},
		{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := SolvePacketEquation(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
