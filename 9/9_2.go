package main

import "sort"

func (p Point) Equals(p2 Point) bool {
	return p[0] == p2[0] && p[1] == p2[1]
}

func GetBasins(input string) int {
	heightMap := ParseInput(input)

	lowPoints := heightMap.FindLowPoints()

	basins := []int{}

	for _, lowPoint := range lowPoints {
		basins = append(basins, heightMap.GetBasinSize(lowPoint))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0] * basins[1] * basins[2]
}

func (hm HeightMap) GetBasinSize(lowPoint Point) int {
	// List of Points
	points := []Point{
		lowPoint,
	}

	previousLength := 0
	// While List of Points changes
	for previousLength < len(points) {
		previousLength = len(points)

		// Iterate over Points in Basin
		for _, point := range points {

			adjacentPoints := hm.GetAdjacentPoints(point)
			for _, adjacentPoint := range adjacentPoints {
				// Check for adjacent point that are !=9 and not in existing list
				// add them to the list
				if hm.Value(adjacentPoint) != 9 && !HasPoint(points, adjacentPoint) {
					points = append(points, adjacentPoint)
				}
			}

		}
	}

	return len(points)
}

func HasPoint(points []Point, new Point) bool {
	for _, point := range points {
		if point.Equals(new) {
			return true
		}
	}

	return false
}
