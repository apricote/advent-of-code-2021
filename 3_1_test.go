package main

import "testing"

func TestGetPowerConsumption(t *testing.T) {
	type test struct {
		input string
		want  uint64
	}

	tests := []test{
		{input: GetDay3ExampleInput(), want: 198},
		{input: GetDay3Input(), want: 2967914},
	}

	for _, tc := range tests {
		got := GetPowerConsumption(tc.input)

		if tc.want != got {
			t.Errorf("Expected power consumption %d but got %d", tc.want, got)
		}
	}
}
