package main

import "testing"

func TestGetOverlappingVents(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay5ExampleInput(), want: 5},
		{input: GetDay5Input(), want: 6397},
	}

	for _, tc := range tests {
		got := GetOverlappingVents(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of overlapping vents to be %d but got %d", tc.want, got)
		}
	}
}
