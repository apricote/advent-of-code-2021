package main

import (
	"strings"
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
	//"github.com/apricote/advent-of-code-2021/util"
)

func TestGetEightLetterCode(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		//{input: util.GetExampleInput(), want: 0},
		{input: util.GetInput(), want: `
###...##..####.###...##..####..##..###.
#..#.#..#....#.#..#.#..#.#....#..#.#..#
###..#......#..#..#.#....###..#..#.###.
#..#.#.....#...###..#....#....####.#..#
#..#.#..#.#....#.#..#..#.#....#..#.#..#
###...##..####.#..#..##..####.#..#.###.
`},
	}

	for _, tc := range tests {
		got := GetEightLetterCode(tc.input)

		if strings.TrimSpace(tc.want) != got {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}
