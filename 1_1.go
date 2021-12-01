package main

import (
	"log"
	"strconv"
	"strings"
)

func GetMeasurementIncreases(input string) (increases int) {
	measurements := parseDay1Input(input)

	increases = getRawMeasurementIncreases(measurements)

	return
}

func parseDay1Input(input string) []int {
	// Convert to string slice
	measurementsRaw := strings.Split(input, "\n")

	// Convert to integer slice
	measurements := make([]int, 0, len(measurementsRaw))
	for _, value := range measurementsRaw {
		integerValue, err := strconv.Atoi(value)
		if err != nil {
			log.Panicf("GetMeasurementIncreases: Can not convert to integer: %s", value)
		}
		measurements = append(measurements, integerValue)
	}

	return measurements
}

func getRawMeasurementIncreases(measurements []int) (increases int) {
	for i, current := range measurements {
		if i == 0 {
			continue
		}

		previous := measurements[i-1]

		if current > previous {
			increases += 1
		}
	}

	return
}
