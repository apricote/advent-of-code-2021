package main

import (
	"strings"
)

type OceanFloor [][]string

func FindTimeToLand(input string) int {
	oceanFloor := ParseInput(input)

	moves := 1
	steps := 0
	for moves > 0 {
		moves = oceanFloor.Step()
		steps += 1
		//log.Printf("OceanFloor Step %v: %v", steps, oceanFloor)
	}

	return steps
}

func ParseInput(input string) OceanFloor {
	oceanFloor := OceanFloor{}

	for i, line := range strings.Split(input, "\n") {
		oceanFloor = append(oceanFloor, make([]string, len(line)))

		for j, char := range strings.Split(line, "") {
			oceanFloor[i][j] = char
		}
	}

	return oceanFloor
}

func (of OceanFloor) String() string {
	output := "\n"

	for _, line := range of {
		for _, char := range line {
			output += char
		}

		output += "\n"
	}

	return output
}

func (of OceanFloor) Step() int {
	moves := 0

	height := len(of)
	width := len(of[0])

	// East
	movable := [][2]int{}
	for i, line := range of {
		for j, char := range line {
			if char != ">" {
				continue
			}

			if of[i][(j+1)%width] != "." {
				continue
			}

			movable = append(movable, [2]int{i, j})
		}
	}

	moves += len(movable)

	for _, cucumber := range movable {
		i, j := cucumber[0], cucumber[1]

		of[i][j] = "."
		of[i][(j+1)%width] = ">"
	}

	// South
	movable = [][2]int{}
	for i, line := range of {
		for j, char := range line {
			if char != "v" {
				continue
			}

			if of[(i+1)%height][j] != "." {
				continue
			}

			movable = append(movable, [2]int{i, j})
		}
	}
	moves += len(movable)

	for _, cucumber := range movable {
		i, j := cucumber[0], cucumber[1]

		of[i][j] = "."
		of[(i+1)%height][j] = "v"
	}

	return moves
}
