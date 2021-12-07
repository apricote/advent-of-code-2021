package main

import (
	"math"
	"sort"
)

func GetMinimumFuelCostWithRaisingPrices(input string) int {
	horizontalPositions := ParseInput(input)

	cheapestPosition := FindMinimalFuelCost(horizontalPositions)

	return GetFuelCostToPositionWithRaisingPrices(cheapestPosition, horizontalPositions)
}

func FindMinimalFuelCost(horizontalPositions []int) int {
	sort.Ints(horizontalPositions)
	min := horizontalPositions[0]
	max := horizontalPositions[len(horizontalPositions)-1]

	minPosition := -1
	minFuelCost := math.MaxInt

	for i := min; i < max; i++ {
		curFuelCost := GetFuelCostToPositionWithRaisingPrices(i, horizontalPositions)

		if curFuelCost < minFuelCost {
			minFuelCost = curFuelCost
			minPosition = i
		}
	}

	return minPosition
}

func GetFuelCostToPositionWithRaisingPrices(targetPos int, positions []int) (sum int) {
	for _, pos := range positions {
		var dist int
		if pos < targetPos {
			dist = targetPos - pos
		} else {
			dist = pos - targetPos
		}

		for dist > 0 {
			sum += dist
			dist--
		}
	}
	return
}
