package main

import (
	"fmt"
	"log"
)

func SolveCurrentDayWithTwist(input string) int {
	riskLevels, err := ParseInput(input)

	if err != nil {
		log.Fatal(err)
	}

	if len(riskLevels) == 0 {
		log.Fatalf("no entries in riskLevels: %v", riskLevels)
	}

	riskLevels = RepeatMap(riskLevels)

	source := [2]int{0, 0}
	destination := [2]int{len(riskLevels[0]) - 1, len(riskLevels) - 1}

	return Dijkstra(riskLevels, source, destination)
}

func RepeatMap(originalRiskLevels [][]int) [][]int {
	originalXLength := len(originalRiskLevels[0])
	originalYLength := len(originalRiskLevels)

	totalXLength := originalXLength * 5
	totalYLength := originalYLength * 5

	riskLevels := [][]int{}

	for y := 0; y < totalYLength; y++ {

		for x := 0; x < totalXLength; x++ {
			if y == 0 {
				// Ugly hack to initialize 2d slice
				riskLevels = append(riskLevels, make([]int, totalXLength))
			}

			loopX := x / originalXLength
			loopY := y / originalYLength

			originalX := x % originalXLength
			originalY := y % originalYLength

			riskLevels[x][y] = originalRiskLevels[originalX][originalY] + loopX + loopY

			if riskLevels[x][y] > 9 {
				riskLevels[x][y] -= 9
			}

		}

	}

	riskLevelsString := ""
	for _, line := range riskLevels {
		for _, value := range line {
			riskLevelsString += fmt.Sprintf("%v", value)
		}
		riskLevelsString += "\n"
	}

	fmt.Print(riskLevelsString)

	return riskLevels
}
