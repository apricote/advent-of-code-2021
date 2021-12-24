package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/apricote/advent-of-code-2021/util"
)

func TestOrganizeAmphipods(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		//{input: util.GetExampleInput(), want: 12521},
		{input: util.GetInput(), want: 0},
	}

	for _, tc := range tests {
		got := OrganizeAmphipods(tc.input)

		if tc.want != got {
			t.Errorf("Expected %d but got %d", tc.want, got)
		}
	}
}

func TestParseInput(t *testing.T) {
	type test struct {
		input string
		want  Burrow
	}

	tests := []test{
		{input: util.GetExampleInput(), want: Burrow{
			&Amphipod{
				Type:     B,
				Room:     RoomA,
				Position: 0,
			},
			&Amphipod{
				Type:     A,
				Room:     RoomA,
				Position: 1,
			},
			&Amphipod{
				Type:     C,
				Room:     RoomB,
				Position: 0,
			},
			&Amphipod{
				Type:     D,
				Room:     RoomB,
				Position: 1,
			},
			&Amphipod{
				Type:     B,
				Room:     RoomC,
				Position: 0,
			},
			&Amphipod{
				Type:     C,
				Room:     RoomC,
				Position: 1,
			},
			&Amphipod{
				Type:     D,
				Room:     RoomD,
				Position: 0,
			},
			&Amphipod{
				Type:     A,
				Room:     RoomD,
				Position: 1,
			},
		}},
		{input: strings.TrimSpace(`
#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
`), want: Burrow{
			&Amphipod{
				Type:     A,
				Room:     RoomA,
				Position: 0,
			},
			&Amphipod{
				Type:     A,
				Room:     RoomA,
				Position: 1,
			},
			&Amphipod{
				Type:     B,
				Room:     RoomB,
				Position: 0,
			},
			&Amphipod{
				Type:     B,
				Room:     RoomB,
				Position: 1,
			},
			&Amphipod{
				Type:     C,
				Room:     RoomC,
				Position: 0,
			},
			&Amphipod{
				Type:     C,
				Room:     RoomC,
				Position: 1,
			},
			&Amphipod{
				Type:     D,
				Room:     RoomD,
				Position: 0,
			},
			&Amphipod{
				Type:     D,
				Room:     RoomD,
				Position: 1,
			},
		}},
	}

	for _, tc := range tests {
		got := ParseInput(tc.input)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Expected %v but got %v", tc.want, got)
		}
	}
}

func TestBurrowIsFinished(t *testing.T) {
	burrow := ParseInput(strings.TrimSpace(`
#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
`))

	if !burrow.IsFinished() {
		t.Errorf("Expected burrow to be finished")
	}
}
