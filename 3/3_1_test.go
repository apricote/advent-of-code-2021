package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestGetPowerConsumption(t *testing.T) {
	type test struct {
		input string
		want  uint64
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 198},
		{input: util.GetInput(), want: 2967914},
	}

	for _, tc := range tests {
		got := GetPowerConsumption(tc.input)

		if tc.want != got {
			t.Errorf("Expected power consumption %d but got %d", tc.want, got)
		}
	}
}
