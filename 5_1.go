package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Line struct {
	Start Coord2D
	End   Coord2D
}

type Coord2D struct {
	X int
	Y int
}

func (c Coord2D) String() string {
	return fmt.Sprintf("%v,%v", c.X, c.Y)
}

func Get2DDistance(a, b Coord2D) float64 {
	return math.Sqrt(
		math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2),
	)
}

func GetOverlappingVents(input string) int {
	lineSegments := ParseDay5Input(input)

	nonDiagonalLineSegments := GetNonDiagonalLineSegment(lineSegments)

	linePoints := GetAllPointsInLines(nonDiagonalLineSegments)

	linePointFrequency := make(map[string]int)
	for _, linePoint := range linePoints {

		if frequency, present := linePointFrequency[linePoint.String()]; present {
			linePointFrequency[linePoint.String()] = frequency + 1
		} else {
			linePointFrequency[linePoint.String()] = 1
		}
	}

	highRiskPoints := 0
	for _, pointFrequency := range linePointFrequency {
		if pointFrequency > 1 {
			highRiskPoints += 1
		}
	}

	return highRiskPoints
}

func GetAllPointsInLines(lineSegments []Line) []Coord2D {
	linePoints := make([]Coord2D, 0, len(lineSegments))

	for _, lineSegment := range lineSegments {
		lineLength := Get2DDistance(lineSegment.Start, lineSegment.End)

		lineVector := Coord2D{
			X: lineSegment.End.X - lineSegment.Start.X,
			Y: lineSegment.End.Y - lineSegment.Start.Y,
		}

		for i := 0.0; i <= lineLength; i++ {
			traveledDistance := i / lineLength

			newPoint := Coord2D{
				X: lineSegment.Start.X + int(math.Round(traveledDistance*float64(lineVector.X))),
				Y: lineSegment.Start.Y + int(math.Round(traveledDistance*float64(lineVector.Y))),
			}

			linePoints = append(linePoints, newPoint)
		}
	}

	return linePoints
}

func GetNonDiagonalLineSegment(lineSegments []Line) []Line {
	nonDiagonalLineSegments := make([]Line, 0, len(lineSegments))
	for _, lineSegment := range lineSegments {
		if lineSegment.Start.X == lineSegment.End.X || lineSegment.Start.Y == lineSegment.End.Y {
			nonDiagonalLineSegments = append(nonDiagonalLineSegments, lineSegment)
		}
	}
	return nonDiagonalLineSegments
}

func ParseDay5Input(input string) []Line {
	inputLines := strings.Split(input, "\n")

	lines := make([]Line, 0, len(inputLines))
	for _, inputLine := range inputLines {
		lines = append(lines, ParseDay5Line(inputLine))
	}

	return lines
}

func ParseDay5Line(input string) Line {
	stringCoordinates := strings.Split(input, " -> ")

	return Line{
		Start: ParseCoord2D(stringCoordinates[0]),
		End:   ParseCoord2D(stringCoordinates[1]),
	}
}

func ParseCoord2D(input string) Coord2D {
	stringCoordinates := strings.Split(input, ",")

	coordinates := make([]int, 0, len(stringCoordinates))

	for _, stringCoord := range stringCoordinates {
		coord, err := strconv.Atoi(stringCoord)
		if err != nil {
			log.Fatalf("Failed to parse coordinate %v", stringCoord)
		}

		coordinates = append(coordinates, coord)
	}

	return Coord2D{
		X: coordinates[0],
		Y: coordinates[1],
	}
}
