package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	maxXVelocity = 1000
	maxYVelocity = 1000
	minXVelocity = -1000
	minYVelocity = -1000
)

// TargetArea contains two ranges of coordinates
type TargetArea [2][2]int

func FindVelocityWithStyle(input string) int {
	targetArea, err := ParseInput(input)
	if err != nil {
		log.Fatalf("unable to parse input: %v", err)
	}

	totalMaxHeight := math.MinInt

	for x := minXVelocity; x <= maxXVelocity; x++ {
		for y := minYVelocity; y <= maxYVelocity; y++ {
			if x > targetArea[1][0] {
				continue
			}

			hits, maxHeight := targetArea.Hits(&[2]int{x, y})

			if hits && maxHeight > totalMaxHeight {
				totalMaxHeight = maxHeight
			}
		}
	}

	return totalMaxHeight
}

func ParseInput(input string) (TargetArea, error) {
	// target area: x=20..30, y=-10..-5
	input = strings.ReplaceAll(input, "target area: ", "")

	// x=20..30, y=-10..-5
	coordRanges := strings.Split(input, ", ")

	// ["x=20..30", "y=-10..-5"]
	coordRanges[0] = strings.ReplaceAll(coordRanges[0], "x=", "")
	coordRanges[1] = strings.ReplaceAll(coordRanges[1], "y=", "")

	// ["20..30", "-10..-5"]
	xCoords := strings.Split(coordRanges[0], "..")
	yCoords := strings.Split(coordRanges[1], "..")

	// [["20","30"], ["-10","-5"]]
	x1, err := strconv.Atoi(xCoords[0])
	if err != nil {
		return TargetArea{}, err
	}
	x2, err := strconv.Atoi(xCoords[1])
	if err != nil {
		return TargetArea{}, err
	}
	y1, err := strconv.Atoi(yCoords[0])
	if err != nil {
		return TargetArea{}, err
	}
	y2, err := strconv.Atoi(yCoords[1])
	if err != nil {
		return TargetArea{}, err
	}

	return TargetArea{
		{x1, y1},
		{x2, y2},
	}, nil
}

func (ta *TargetArea) Hits(velocity *[2]int) (bool, int) {
	position := [2]int{0, 0}

	maxHeight := position[1]

	for {
		position[0] += velocity[0]
		position[1] += velocity[1]

		if velocity[0] > 0 {
			velocity[0] -= 1
		} else if velocity[0] < 0 {
			velocity[0] += 1
		}

		velocity[1] -= 1

		// Check maxHeight
		if position[1] > maxHeight {
			maxHeight = position[1]
		}

		// Check if in TargetArea
		if ta.IsIn(&position) {
			return true, maxHeight
		}

		// Check if we can abort the search
		// velocity is negativ and below target area
		if velocity[1] < 0 && position[1] < ta[0][1] {
			break
		}
	}

	return false, 0
}

func (ta *TargetArea) IsIn(coords *[2]int) bool {
	return ta[0][0] <= coords[0] && // x1 <= x
		coords[0] <= ta[1][0] && // x <= x2
		ta[0][1] <= coords[1] && // y1 <= y
		coords[1] <= ta[1][1] // y <= y2
}
