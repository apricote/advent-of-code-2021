package main

import "testing"

func TestGetOverlappingVentsWithDiagonals(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay5ExampleInput(), want: 12},
		{input: GetDay5Input(), want: 22335},
	}

	for _, tc := range tests {
		got := GetOverlappingVentsWithDiagonals(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of overlapping vents (with diagonals) to be %d but got %d", tc.want, got)
		}
	}
}
