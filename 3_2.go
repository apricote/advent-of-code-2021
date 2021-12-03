package main

import (
	"log"
	"strconv"
	"strings"
)

func GetLifeSupportRating(input string) int64 {
	inputRows := strings.Split(input, "\n")

	oxygenGeneratorRating := GetMachineRating(inputRows, BitOn, BitOff)
	co2ScrubberRating := GetMachineRating(inputRows, BitOff, BitOn)

	return oxygenGeneratorRating * co2ScrubberRating
}

func GetMachineRating(inputRows []string, bitOne, bitTwo rune) int64 {
	inputBits := len(inputRows[0])

	remainingNumbers := make([]string, len(inputRows))
	copy(remainingNumbers, inputRows)

	for i := 0; i < inputBits && len(remainingNumbers) > 1; i++ {
		var oneCount int

		for _, number := range remainingNumbers {
			if string(BitOn) == strings.Split(number, "")[i] {
				oneCount += 1
			}
		}

		var mostCommonValue rune

		if float64(oneCount) >= float64(len(remainingNumbers))/2 {
			mostCommonValue = bitOne
		} else {
			mostCommonValue = bitTwo
		}

		newRemainingNumbers := make([]string, 0, len(remainingNumbers))

		for _, number := range remainingNumbers {
			if string(mostCommonValue) == strings.Split(number, "")[i] {
				newRemainingNumbers = append(newRemainingNumbers, number)
			}
		}

		remainingNumbers = newRemainingNumbers
	}

	machineRating, err := strconv.ParseInt(remainingNumbers[0], 2, 0)
	if err != nil {
		log.Panicf("unable to convert to uint: %s", remainingNumbers[0])
	}

	return machineRating
}
