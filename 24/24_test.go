package main

import (
	"reflect"
	"testing"
)

func TestBacktrack(t *testing.T) {
	highestModelNumber, rejected := Backtrack([]int{}, true)

	if rejected {
		t.Errorf("Expected Backtrack not to be rejected")
	}
	if !reflect.DeepEqual(highestModelNumber, []int{4, 9, 9, 1, 7, 9, 2, 9, 9, 3, 4, 9, 9, 9}) {
		t.Errorf("Expected highest model number but got %v", highestModelNumber)
	}

	lowestModelNumber, rejected := Backtrack([]int{}, false)

	if rejected {
		t.Errorf("Expected Backtrack not to be rejected")
	}
	if !reflect.DeepEqual(lowestModelNumber, []int{1, 1, 9, 1, 1, 3, 1, 6, 7, 1, 1, 8, 1, 6}) {
		t.Errorf("Expected lowest model number but got %v", lowestModelNumber)
	}
}
