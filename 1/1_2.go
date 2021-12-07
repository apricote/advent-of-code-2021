package main

func GetSlidingWindowMeasurementIncreases(input string) (increases int) {
	measurements := parseDay1Input(input)

	// Make sliding window slice
	slidingWindowMeasurements := make([]int, 0, len(measurements))
	for i := range measurements {
		if i >= len(measurements)-2 {
			// impossible to make three-measurement sum with only 2 elements left
			break
		}

		sum := measurements[i] + measurements[i+1] + measurements[i+2]

		slidingWindowMeasurements = append(slidingWindowMeasurements, sum)
	}

	increases = getRawMeasurementIncreases(slidingWindowMeasurements)
	return
}
