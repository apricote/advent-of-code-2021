package main

import (
	"log"
	"math"
	"strings"
)

func GetEightLetterCode(inputString string) string {
	input, err := ParseInput(inputString)
	if err != nil {
		log.Fatal(err)
	}

	paper := input.Paper

	for _, fold := range input.Folds {
		paper = paper.Fold(fold)
	}

	log.Printf("Paper: \n%v", paper)

	return paper.String()
}

func (p Paper) String() string {
	maxX := math.MinInt
	maxY := math.MinInt

	for _, dot := range p {
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
	}

	output := ""

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			symbol := "."
			if p.Includes([2]int{x, y}) {
				symbol = "#"
			}

			output += symbol
		}

		output += "\n"
	}

	return strings.TrimSpace(output)
}
