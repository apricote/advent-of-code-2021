package main

import "testing"

func TestCountLanternfish256(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay6ExampleInput(), want: 26984457539},
		{input: GetDay6Input(), want: 1601616884019},
	}

	for _, tc := range tests {
		got := CountLanternfish256(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of lanternfishes after 256 days to be %d but got %d", tc.want, got)
		}
	}
}
