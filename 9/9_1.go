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

	lowPoints := heightMap.FindLowPoints()

	totalRiskLevel := 0

	for _, lowPoint := range lowPoints {
		totalRiskLevel += 1 + heightMap.Value(lowPoint)
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

func (hm HeightMap) FindLowPoints() []Point {
	lowPoints := []Point{}

	for i, line := range hm {
		for j := range line {
			point := Point{i, j}
			if hm.IsLowPoint(point) {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return lowPoints
}

func (hm HeightMap) IsLowPoint(point Point) bool {
	value := hm[point[0]][point[1]]

	adjacentPoints := hm.GetAdjacentPoints(point)

	for _, adjacentPoint := range adjacentPoints {
		if value >= hm.Value(adjacentPoint) {
			return false
		}
	}

	return true
}

func (hm HeightMap) GetAdjacentPoints(point Point) []Point {
	adjacentPoints := []Point{}

	// Top
	if point[0] > 0 {
		adjacentPoints = append(adjacentPoints, Point{
			point[0] - 1, point[1],
		})
	}
	// Bottom
	if point[0] < len(hm)-1 {
		adjacentPoints = append(adjacentPoints, Point{
			point[0] + 1, point[1],
		})
	}
	// Left
	if point[1] > 0 {
		adjacentPoints = append(adjacentPoints, Point{
			point[0], point[1] - 1,
		})
	}
	// Right
	if point[1] < len(hm[point[0]])-1 {
		adjacentPoints = append(adjacentPoints, Point{
			point[0], point[1] + 1,
		})
	}

	return adjacentPoints
}

func (hm HeightMap) Value(p Point) int {
	return hm[p[0]][p[1]]
}
