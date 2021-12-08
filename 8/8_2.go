package main

import (
	"math"
	"math/bits"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func GetOutputSum(input string) int {
	inputElements := ParseInput(input)

	var totalOutputValue int

	for _, inputElement := range inputElements {
		patternToDigits := DeduceDigits(inputElement.UniqueSignalPatterns)

		outputValue := 0

		for i, outputPattern := range inputElement.OutputValue {
			// i=0 *1000 *10^(3-i)
			// i=1 *100

			digit := patternToDigits[outputPattern]

			outputValue += digit * (powInt(10, 3-i))

		}

		totalOutputValue += outputValue
	}

	return totalOutputValue
}

func DeduceDigits(patterns [10]SevenSegmentOutput) map[SevenSegmentOutput]int {
	digits := map[int]SevenSegmentOutput{}

	for _, pattern := range patterns {
		onesCount := bits.OnesCount8(pattern)

		switch onesCount {
		case 2:
			// unsolvedPattern is 1
			// 1 => CF

			digits[1] = pattern

			continue
		case 4:
			// unsolvedPattern is 4
			// 4 => BCDF

			digits[4] = pattern

			continue
		case 3:
			// unsolvedPattern is 7
			// 7 => ACF

			digits[7] = pattern

			continue
		case 7:
			// unsolvedPattern is 8
			// ABCDEFG

			digits[8] = pattern

			// we can not get any useful information from this!
			continue
		default:
			// Only those 4 cases have unique bitcounts
		}

	}

	// Length 5 => 2,3,5
	patternsLength5 := make([]SevenSegmentOutput, 0)
	for _, pattern := range patterns {
		if bits.OnesCount8(pattern) == 5 {
			patternsLength5 = append(patternsLength5, pattern)
		}
	}
	for _, length5 := range patternsLength5 {
		// 00100100
		//&01110011
		//=00100000
		// ^ Counterexample

		// 00100100
		//&01110111
		//=00100100
		// ^ Example for Digit "3"

		if bits.OnesCount8(digits[1]&length5) == 2 {
			// length5 is 3
			// ACDFG

			digits[3] = length5
		} else if bits.OnesCount8(digits[4]&length5) == 2 {
			// length5 is 2
			// ACDEG

			digits[2] = length5

		} else if bits.OnesCount8(digits[4]&length5) == 3 {
			// length5 is 5
			// ABDFG
			digits[5] = length5
		}
	}

	// Length 6 => 0,6,9
	patternsLength6 := make([]SevenSegmentOutput, 0)
	for _, pattern := range patterns {
		if bits.OnesCount8(pattern) == 6 {
			patternsLength6 = append(patternsLength6, pattern)
		}
	}
	for _, length6 := range patternsLength6 {
		if bits.OnesCount8(digits[1]&length6) == 1 {
			// length6 is 6
			// ABDFG

			digits[6] = length6
		} else if bits.OnesCount8(digits[3]&length6) == 4 {
			// length6 is 0
			// ABCEFG

			digits[0] = length6
		} else if bits.OnesCount8(digits[3]&length6) == 5 {
			// length6 is 9
			// ABCEFG

			digits[9] = length6
		}
	}

	patternToDigits := map[SevenSegmentOutput]int{}
	for digit, pattern := range digits {
		patternToDigits[pattern] = digit
	}

	return patternToDigits
}
