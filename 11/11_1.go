package main

import (
	"log"
	"strconv"
	"strings"
)

type EnergyMap [][]int

// Coordinates for octopus in EnergyMap
type Octopus [2]int

const (
	EMSize = 10
)

func CountFlashesIn100Steps(input string) int {
	energyMap := ParseInput(input)

	count := 0
	for i := 0; i < 100; i++ {
		count += energyMap.SimulateStep()
	}

	return count
}

func ParseInput(input string) EnergyMap {
	energyMap := EnergyMap{}

	for _, line := range strings.Split(input, "\n") {
		levelsInLine := []int{}

		for _, levelString := range strings.Split(line, "") {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatalf("unable to parse energy level %v", levelString)
			}

			levelsInLine = append(levelsInLine, level)
		}

		energyMap = append(energyMap, levelsInLine)
	}

	return energyMap
}

func (em EnergyMap) SimulateStep() (flashes int) {
	// 1. Increase energy level by one
	em.increaseEnergyLevel()

	// 2. Flashes
	// {{0,0},{0,1}}
	alreadyFlashed := []Octopus{}

	var flashIfNecessary func(octoput Octopus)
	flashIfNecessary = func(octopus Octopus) {
		level := em[octopus[0]][octopus[1]]

		if level < 10 {
			return
		}

		if ContainsOctopus(alreadyFlashed, octopus) {
			return
		}

		// FLASH!
		alreadyFlashed = append(alreadyFlashed, octopus)

		// Get Adjacent Octopi
		adjacentOctopi := em.GetAdjacentOctopi(octopus)

		for _, adjacentOctopus := range adjacentOctopi {
			// Raise their energy levels
			em[adjacentOctopus[0]][adjacentOctopus[1]] += 1
			// FLASH!
			flashIfNecessary(adjacentOctopus)
		}
	}

	for i, line := range em {
		for j := range line {
			octopus := Octopus{i, j}
			flashIfNecessary(octopus)
		}
	}

	// 3. Reset to 0
	em.resetSpentEnergy()

	return len(alreadyFlashed)
}

func (em EnergyMap) increaseEnergyLevel() {
	for i, line := range em {
		for j := range line {
			em[i][j] += 1
		}
	}
}

func (em EnergyMap) resetSpentEnergy() {
	for i, line := range em {
		for j, level := range line {
			if level > 9 {
				em[i][j] = 0
			}
		}
	}
}

func ContainsOctopus(octopusList []Octopus, octopus Octopus) bool {
	for _, octopusB := range octopusList {
		if octopus[0] == octopusB[0] && octopus[1] == octopusB[1] {
			return true
		}
	}

	return false
}

func (em EnergyMap) GetAdjacentOctopi(octopus Octopus) []Octopus {
	octopi := []Octopus{}

	row := octopus[0]
	col := octopus[1]

	notTop := row > 0
	notBottom := row < EMSize-1
	notLeft := col > 0
	notRight := col < EMSize-1

	// Top
	if notTop {
		// Top Left
		if notLeft {
			octopi = append(octopi, Octopus{row - 1, col - 1})
		}

		// Top Top
		octopi = append(octopi, Octopus{row - 1, col})

		// Top Right

		if notRight {
			octopi = append(octopi, Octopus{row - 1, col + 1})
		}
	}

	// Left Middle
	if notLeft {
		octopi = append(octopi, Octopus{row, col - 1})
	}

	// Right Middle
	if notRight {
		octopi = append(octopi, Octopus{row, col + 1})
	}

	// Bottom
	if notBottom {
		// Bottom Left
		if notLeft {
			octopi = append(octopi, Octopus{row + 1, col - 1})
		}

		// Bottom Bottom
		octopi = append(octopi, Octopus{row + 1, col})

		// Bottom Right
		if notRight {
			octopi = append(octopi, Octopus{row + 1, col + 1})
		}
	}

	return octopi
}
