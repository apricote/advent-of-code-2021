package main

import "testing"

func TestGetLifeSupportRating(t *testing.T) {
	type test struct {
		input string
		want  int64
	}

	tests := []test{
		{input: GetDay3ExampleInput(), want: 230},
		{input: GetDay3Input(), want: 7041258},
	}

	for _, tc := range tests {
		got := GetLifeSupportRating(tc.input)

		if tc.want != got {
			t.Errorf("Expected life support rating %d but got %d", tc.want, got)
		}
	}
}
