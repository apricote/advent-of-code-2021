package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	LanternfishCycle  = 6
	LanternfishGrowUp = 2
)

func CountLanternfish(input string) int {
	ages := ParseDay6Input(input)

	for i := 0; i < 80; i++ {
		ages = SimulateLanternfishes(ages)
	}

	return CountAllLanternfish(ages)
}

func CountAllLanternfish(ages [9]int) int {
	totalCount := 0
	for _, count := range ages {
		totalCount += count
	}

	return totalCount
}

func ParseDay6Input(input string) [9]int {
	lanterfishAgeStrings := strings.Split(input, ",")

	ages := [9]int{}

	for _, ageString := range lanterfishAgeStrings {
		age, err := strconv.Atoi(ageString)
		if err != nil {
			log.Fatalf("Unable to parse lanternfish age: %v", ageString)
		}

		ages[age]++
	}

	return ages
}

func SimulateLanternfishes(ages [9]int) [9]int {
	newAges := [9]int{}

	for i := 0; i < 9; i++ {
		count := ages[i]

		if i == 0 {
			newAges[LanternfishCycle] += count
			newAges[LanternfishCycle+LanternfishGrowUp] += count
		} else {
			newAges[i-1] += count
		}
	}

	return newAges
}
