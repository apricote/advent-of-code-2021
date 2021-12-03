package main

import "testing"

func TestGetDivePositionWithAim(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay2ExampleInput(), want: 900},
		{input: GetDay2Input(), want: 1942068080},
	}

	for _, tc := range tests {
		got := GetDivePositionWithAim(tc.input)

		if tc.want != got {
			t.Errorf("Expected dive position %d but got %d", tc.want, got)
		}
	}
}
