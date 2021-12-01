package main

import (
	"log"
	"strconv"
	"strings"
)

func GetMeasurementIncreases(input string) (increases int) {
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
