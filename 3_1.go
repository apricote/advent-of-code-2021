package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

const (
	BitOn  = '1'
	BitOff = '0'
)

func GetPowerConsumption(input string) uint64 {
	inputRows := strings.Split(input, "\n")

	countOnes := CountOnes(inputRows)

	var gammaBuffer bytes.Buffer
	var epsilonBuffer bytes.Buffer

	for _, countOne := range countOnes {
		if countOne >= len(inputRows)/2 {
			gammaBuffer.WriteRune(BitOn)
			epsilonBuffer.WriteRune(BitOff)
		} else {
			gammaBuffer.WriteRune(BitOff)
			epsilonBuffer.WriteRune(BitOn)
		}
	}

	gamma, err := strconv.ParseUint(gammaBuffer.String(), 2, 64)
	if err != nil {
		log.Panicf("unable to convert to uint: %s", gammaBuffer.String())
	}

	epsilon, err := strconv.ParseUint(epsilonBuffer.String(), 2, 64)
	if err != nil {
		log.Panicf("unable to convert to uint: %s", epsilonBuffer.String())
	}

	return gamma * epsilon
}

func CountOnes(inputRows []string) []int {
	// Only works because or input is only "0" and "1"
	// Breaks with utf-8
	diagnosticReportLength := len(inputRows[0])

	countOnes := make([]int, diagnosticReportLength)

	for _, inputRow := range inputRows {
		for i, digit := range inputRow {
			if digit == BitOn {
				countOnes[i] += 1
			}
		}
	}

	return countOnes
}
