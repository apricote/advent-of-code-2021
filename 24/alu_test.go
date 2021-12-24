package main

import (
	"strings"
	"testing"
)

func TestRunALU(t *testing.T) {
	type test struct {
		program string
		inputs  []int
		want    [4]int
	}

	tests := []test{
		{
			program: strings.TrimSpace(`
inp x
mul x -1
		`),
			inputs: []int{4},
			want:   [4]int{0, -4, 0, 0},
		},
		{
			program: strings.TrimSpace(`
inp x
mul x -1
		`),
			inputs: []int{9},
			want:   [4]int{0, -9, 0, 0},
		},
		{
			program: strings.TrimSpace(`
inp z
inp x
mul z 3
eql z x
		`),
			inputs: []int{3, 9},
			want:   [4]int{0, 9, 0, 1},
		},
		{
			program: strings.TrimSpace(`
inp z
inp x
mul z 3
eql z x
		`),
			inputs: []int{4, 9},
			want:   [4]int{0, 9, 0, 0},
		},
		{
			program: strings.TrimSpace(`
inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2
		`),
			inputs: []int{13},
			want:   [4]int{1, 1, 0, 1},
		},
	}

	for _, tc := range tests {
		got := RunALU(tc.program, tc.inputs)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}
