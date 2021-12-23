package main

import (
	"testing"
)

func TestParseSnailfishNumber(t *testing.T) {
	type test struct {
		input string
		want  SnailfishNumber
	}

	tests := []test{
		{input: "[1,2]", want: SnailfishNumberPair{1, 2}},
		{input: "[[1,2],3]", want: SnailfishNumberPair{SnailfishNumberPair{1, 2}, 3}},
		{input: "[9,[8,7]]", want: SnailfishNumberPair{9, SnailfishNumberPair{8, 7}}},
		{input: "[[1,9],[8,5]]", want: SnailfishNumberPair{SnailfishNumberPair{1, 9}, SnailfishNumberPair{8, 5}}},
		{
			input: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
			want: SnailfishNumberPair{
				SnailfishNumberPair{
					SnailfishNumberPair{
						SnailfishNumberPair{1, 2},
						SnailfishNumberPair{3, 4}},
					SnailfishNumberPair{
						SnailfishNumberPair{5, 6},
						SnailfishNumberPair{7, 8}}},
				9},
		},
		{
			input: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
			want: SnailfishNumberPair{
				SnailfishNumberPair{
					SnailfishNumberPair{
						SnailfishNumberPair{1, 3},
						SnailfishNumberPair{5, 3}},
					SnailfishNumberPair{
						SnailfishNumberPair{1, 3},
						SnailfishNumberPair{8, 7}}},
				SnailfishNumberPair{
					SnailfishNumberPair{
						SnailfishNumberPair{4, 9},
						SnailfishNumberPair{6, 9}},
					SnailfishNumberPair{
						SnailfishNumberPair{8, 2},
						SnailfishNumberPair{7, 3}}}},
		},
	}

	for _, tc := range tests {
		got, err := ParseSnailfishNumber(tc.input)
		if err != nil {
			t.Error(err)
		}

		if tc.want != got {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}
