package main

import "testing"

func TestGetMeasurementIncreases(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: GetDay1ExampleInput(), want: 7},
		{input: GetDay1Input(), want: 1233},
	}

	for _, tc := range tests {
		got := GetMeasurementIncreases(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d increases but got %d", tc.want, got)
		}
	}
}
