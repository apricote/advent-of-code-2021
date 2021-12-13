package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Paper [][2]int

type Fold struct {
	Axis  Axis
	Value int
}

type Axis string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
)

type Input struct {
	Paper Paper
	Folds []Fold
}

func CountDotsAfterOneFold(inputString string) int {
	input, err := ParseInput(inputString)
	if err != nil {
		log.Fatal(err)
	}

	// Do one fold
	if len(input.Folds) < 1 {
		log.Fatal("no folds available!")
	}

	paper := input.Paper.Fold(input.Folds[0])

	return len(paper)
}

func ParseInput(input string) (Input, error) {
	parts := strings.Split(input, "\n\n")

	if len(parts) != 2 {
		return Input{}, fmt.Errorf("found more than two parts in instructions: %v", input)
	}

	coordinates := parts[0]
	foldInstructions := parts[1]

	// Parsing the Paper
	paper := Paper{}

	for _, dot := range strings.Split(coordinates, "\n") {
		xy := strings.Split(dot, ",")

		if len(xy) != 2 {

			return Input{}, fmt.Errorf("found more than two elements in line: %v", dot)
		}

		x, err := strconv.Atoi(xy[0])
		if err != nil {
			return Input{}, fmt.Errorf("unable to parse x coordinate: %v", xy[0])
		}

		y, err := strconv.Atoi(xy[1])
		if err != nil {
			return Input{}, fmt.Errorf("unable to parse y coordinate: %v", xy[1])
		}

		paper = append(paper, [2]int{x, y})
	}

	// Parsing the fold instructions
	folds := []Fold{}

	for _, foldInstruction := range strings.Split(foldInstructions, "\n") {
		foldInstruction = strings.ReplaceAll(foldInstruction, "fold along ", "")

		instructions := strings.Split(foldInstruction, "=")

		if len(instructions) != 2 {
			return Input{}, fmt.Errorf("unable to parse fold instructions, more/less than 2 elements: %v", foldInstruction)
		}

		var axis Axis

		switch instructions[0] {
		case "x":
			axis = AxisX
		case "y":
			axis = AxisY
		}

		value, err := strconv.Atoi(instructions[1])
		if err != nil {
			return Input{}, fmt.Errorf("unable to parse fold value: %v", instructions[1])
		}

		folds = append(folds, Fold{
			Axis:  axis,
			Value: value,
		})
	}

	return Input{
		Paper: paper,
		Folds: folds,
	}, nil
}

func (p Paper) Fold(fold Fold) Paper {
	newPaper := Paper{}

	for _, dot := range p {
		x := dot[0]
		y := dot[1]
		// 6,0 -> 4,0
		// 9,0 -> 1,0
		if fold.Axis == AxisX && x > fold.Value {
			x = fold.Value - (x - fold.Value)
		}

		if fold.Axis == AxisY && y > fold.Value {
			y = fold.Value - (y - fold.Value)
		}

		newDot := [2]int{x, y}

		if !newPaper.Includes(newDot) {
			newPaper = append(newPaper, newDot)
		}
	}

	return newPaper
}

func (p Paper) Includes(coords [2]int) bool {
	for _, dot := range p {
		if dot[0] == coords[0] && dot[1] == coords[1] {
			return true
		}
	}

	return false
}
