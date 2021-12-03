package main

import "testing"

func TestGetDivePosition(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay2ExampleInput(), want: 150},
		{input: GetDay2Input(), want: 2039912},
	}

	for _, tc := range tests {
		got := GetDivePosition(tc.input)

		if tc.want != got {
			t.Errorf("Expected dive position %d but got %d", tc.want, got)
		}
	}
}
