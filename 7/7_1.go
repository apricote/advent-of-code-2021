package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetMinimumFuelCost(input string) int {
	horizontalPositions := ParseInput(input)

	meanPosition := GetMean(horizontalPositions)

	return GetFuelCostToPosition(meanPosition, horizontalPositions)
}

func ParseInput(input string) []int {
	positionStrings := strings.Split(input, ",")

	horizontalPositions := make([]int, 0, len(positionStrings))
	for _, posString := range positionStrings {
		pos, err := strconv.Atoi(posString)
		if err != nil {
			log.Fatalf("unable to parse horizontal position %v", posString)
		}
		horizontalPositions = append(horizontalPositions, pos)
	}

	return horizontalPositions
}

func GetFuelCostToPosition(targetPos int, positions []int) (sum int) {
	for _, pos := range positions {
		if pos < targetPos {
			sum += targetPos - pos
		} else {
			sum += pos - targetPos
		}
	}
	return
}

func GetMean(slice []int) int {
	sort.Ints(slice)
	length := len(slice)

	if length%2 == 1 {
		return slice[length/2]
	} else {
		a := slice[length/2]
		b := slice[length/2-1]

		return (a + b) / 2
	}
}
