package main

func GetOverlappingVentsWithDiagonals(input string) int {
	lineSegments := ParseDay5Input(input)

	linePoints := GetAllPointsInLines(lineSegments)

	return CalculateHighRiskPoints(linePoints)
}
