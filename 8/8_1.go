package main

import (
	"math/bits"
	"strings"
)

// 0gfedcba
type SevenSegmentOutput = uint8

const (
	SegmentA SevenSegmentOutput = 1 << iota
	SegmentB SevenSegmentOutput = 1 << iota
	SegmentC SevenSegmentOutput = 1 << iota
	SegmentD SevenSegmentOutput = 1 << iota
	SegmentE SevenSegmentOutput = 1 << iota
	SegmentF SevenSegmentOutput = 1 << iota
	SegmentG SevenSegmentOutput = 1 << iota
)

type InputElement struct {
	UniqueSignalPatterns [10]SevenSegmentOutput
	OutputValue          [4]SevenSegmentOutput
}

func SevenSegmentSearch(input string) int {
	inputElements := ParseInput(input)

	var digitsWithUniqueSegments int

	for _, element := range inputElements {
		for _, digit := range element.OutputValue {
			onesCount := bits.OnesCount8(digit)

			if onesCount == 2 || onesCount == 4 || onesCount == 3 || onesCount == 7 {
				digitsWithUniqueSegments += 1
			}
		}
	}

	return digitsWithUniqueSegments
}

func ParseInput(input string) []InputElement {
	inputElements := make([]InputElement, 0)

	for _, inputLine := range strings.Split(input, "\n") {
		lists := strings.Split(inputLine, " | ")

		uniqueSignalPatterns := ParseSevenSegmentOutputList(lists[0])
		outputValue := ParseSevenSegmentOutputList(lists[1])

		element := InputElement{}

		copy(element.UniqueSignalPatterns[:], uniqueSignalPatterns)
		copy(element.OutputValue[:], outputValue)

		inputElements = append(inputElements, element)
	}

	return inputElements
}

func ParseSevenSegmentOutputList(input string) []SevenSegmentOutput {
	var outputSlice []SevenSegmentOutput

	for _, outputString := range strings.Split(input, " ") {
		outputSlice = append(outputSlice, ParseSevenSegmentOutput(outputString))
	}

	return outputSlice
}

func ParseSevenSegmentOutput(input string) SevenSegmentOutput {
	var output SevenSegmentOutput

	for _, digit := range strings.Split(input, "") {
		switch digit {
		case "a":
			output |= SegmentA
		case "b":
			output |= SegmentB
		case "c":
			output |= SegmentC
		case "d":
			output |= SegmentD
		case "e":
			output |= SegmentE
		case "f":
			output |= SegmentF
		case "g":
			output |= SegmentG
		}
	}

	return output
}
