package main

import (
	"log"
)

func FindAllPossibleVelocities(input string) int {
	targetArea, err := ParseInput(input)
	if err != nil {
		log.Fatalf("unable to parse input: %v", err)
	}

	count := 0

	for x := minXVelocity; x <= maxXVelocity; x++ {
		for y := minYVelocity; y <= maxYVelocity; y++ {
			hits, _ := targetArea.Hits([2]int{x, y})

			if hits {
				count += 1
			}
		}
	}

	return count
}
