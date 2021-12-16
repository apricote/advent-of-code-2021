package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetVersionNumberCountFromBITS(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "8A004A801A8002F478", want: 16},
		{input: "620080001611562C8802118E34", want: 12},
		{input: "C0015000016115A2E0802F182340", want: 23},
		{input: "A0016C880162017C3686B18A3D4780", want: 31},
		{input: util.GetInput(), want: 940},
	}

	for _, tc := range tests {
		got := GetVersionNumberCountFromBITS(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
