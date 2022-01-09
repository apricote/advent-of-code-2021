package main

import (
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
	//"github.com/apricote/advent-of-code-2021/util"
)

func TestEnhanceImageFiftyTimes(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: util.GetExampleInput(), want: 3351},
		{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := EnhanceImageFiftyTimes(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
