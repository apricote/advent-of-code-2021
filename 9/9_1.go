package main

import (
	"log"
	"strconv"
	"strings"
)

type HeightMap [][]int

type Point [2]int

func SumRiskLevels(input string) int {
	heightMap := ParseInput(input)

	lowPoints := FindLowPoints(heightMap)

	totalRiskLevel := 0

	for _, lowPoint := range lowPoints {
		totalRiskLevel += 1 + heightMap[lowPoint[0]][lowPoint[1]]
	}

	return totalRiskLevel
}

func ParseInput(input string) HeightMap {
	heightMap := [][]int{}

	for i, line := range strings.Split(input, "\n") {
		heightMap = append(heightMap, make([]int, len(line)))

		for j, point := range strings.Split(line, "") {
			height, err := strconv.Atoi(point)
			if err != nil {
				log.Fatalf("unable to parse height map value %v", point)
			}

			heightMap[i][j] = height
		}
	}

	return heightMap
}

func FindLowPoints(heightMap HeightMap) []Point {
	lowPoints := []Point{}

	for i, line := range heightMap {
		for j := range line {
			point := Point{i, j}
			if IsLowPoint(heightMap, point) {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return lowPoints
}

func IsLowPoint(heightMap HeightMap, point Point) bool {
	value := heightMap[point[0]][point[1]]

	sourroundingValues := []int{}

	// Top
	if point[0] > 0 {
		sourroundingValues = append(sourroundingValues, heightMap[point[0]-1][point[1]])
	}
	// Bottom
	if point[0] < len(heightMap)-1 {
		sourroundingValues = append(sourroundingValues, heightMap[point[0]+1][point[1]])
	}
	// Left
	if point[1] > 0 {
		sourroundingValues = append(sourroundingValues, heightMap[point[0]][point[1]-1])
	}
	// Right
	if point[1] < len(heightMap[point[0]])-1 {
		sourroundingValues = append(sourroundingValues, heightMap[point[0]][point[1]+1])
	}

	for _, otherValue := range sourroundingValues {
		if value >= otherValue {
			return false
		}
	}

	return true
}
