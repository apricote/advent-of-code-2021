package main

import "testing"

func TestGetSlidingWindowMeasurementIncreases(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay1ExampleInput(), want: 5},
		{input: GetDay1Input(), want: 1275},
	}

	for _, tc := range tests {
		got := GetSlidingWindowMeasurementIncreases(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d increases but got %d", tc.want, got)
		}
	}
}
