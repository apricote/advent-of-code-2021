package main

import "testing"

func TestCountLanternfish(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay6ExampleInput(), want: 5934},
		{input: GetDay6Input(), want: 352151},
	}

	for _, tc := range tests {
		got := CountLanternfish(tc.input)

		if tc.want != got {
			t.Errorf("Expected number of lanternfishes to be %d but got %d", tc.want, got)
		}
	}
}
