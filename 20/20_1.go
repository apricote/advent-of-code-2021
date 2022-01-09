package main

import (
	"log"
	"strings"
)

type Input struct {
	ImageEnhancementAlgorithm [512]bool
	Image                     Image
}

type Image struct {
	Inside  [][]bool
	Outside bool
}

func EnhanceImageTwice(inputString string) int {
	input := ParseInput(inputString)

	image := input.Image
	image = image.Enhance(input.ImageEnhancementAlgorithm)
	image = image.Enhance(input.ImageEnhancementAlgorithm)

	if image.Outside {
		log.Fatalf("can not count number of infinity bits")
	}

	count := 0

	for _, line := range image.Inside {
		for _, bit := range line {
			if bit {
				count += 1
			}
		}
	}

	return count
}

func ParseInput(input string) Input {
	parts := strings.Split(input, "\n\n")

	// Image enhancement algorithm
	imageEnhancementAlgorithm := [512]bool{}
	for i, value := range strings.Split(parts[0], "") {
		imageEnhancementAlgorithm[i] = ParseValue(value)
	}

	// image
	image := [][]bool{}
	for _, line := range strings.Split(parts[1], "\n") {
		imageLine := []bool{}

		for _, value := range strings.Split(line, "") {
			imageLine = append(imageLine, ParseValue(value))
		}

		image = append(image, imageLine)
	}

	return Input{
		ImageEnhancementAlgorithm: imageEnhancementAlgorithm,
		Image: Image{
			Inside:  image,
			Outside: false,
		}}
}

func ParseValue(value string) bool {
	switch value {
	case "#":
		return true
	case ".":
		return false
	}

	return false
}

func (i Image) Enhance(algorithm [512]bool) Image {
	inputInsideSizeX := len(i.Inside[0])
	inputInsideSizeY := len(i.Inside)

	// Initialize Output Inside
	outputInside := [][]bool{}

	for y := 0; y < inputInsideSizeY+2; y++ {
		outputInside = append(outputInside, make([]bool, inputInsideSizeX+2))
	}

	// Calculate new Outside
	outputOutside := false
	switch i.Outside {
	case true:
		outputOutside = algorithm[511]
	case false:
		outputOutside = algorithm[0]
	}

	// Enhance inside of image
	for y, outputLine := range outputInside {
		for x := range outputLine {
			inputNumber := InputSequenceToNumber(i.GetInputSequence(x-1, y-1))
			outputInside[y][x] = algorithm[inputNumber]
		}
	}

	// Return new Image
	return Image{
		Inside:  outputInside,
		Outside: outputOutside,
	}
}

func (i Image) GetInputSequence(x, y int) [9]bool {
	return [9]bool{
		i.GetValue(x-1, y-1),
		i.GetValue(x, y-1),
		i.GetValue(x+1, y-1),
		i.GetValue(x-1, y),
		i.GetValue(x, y),
		i.GetValue(x+1, y),
		i.GetValue(x-1, y+1),
		i.GetValue(x, y+1),
		i.GetValue(x+1, y+1),
	}
}

func (i Image) GetValue(x, y int) bool {
	if x < 0 || y < 0 || x > len(i.Inside[0])-1 || y > len(i.Inside)-1 {
		return i.Outside
	}

	return i.Inside[y][x]
}

func InputSequenceToNumber(inputSequence [9]bool) int {
	number := 0
	for i, bit := range inputSequence {
		if bit {
			number += 1 << (8 - i)
		}
	}

	return number
}
