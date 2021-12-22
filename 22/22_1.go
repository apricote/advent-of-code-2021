package main

import (
	"log"
	"strconv"
	"strings"
)

type Coord3D [3]int
type Command string

type RebootStep struct {
	Command Command
	CoordA  Coord3D
	CoordB  Coord3D
}

const (
	CommandOn  Command = "on"
	CommandOff Command = "off"
)

func CountOnCubes(input string) int {
	steps, err := ParseInput(input)
	if err != nil {
		log.Fatalf("Unable to parse reboot steps: %v", err)
	}

	reactor := map[Coord3D]bool{}

	for _, step := range steps {
		state := false
		switch step.Command {
		case CommandOn:
			state = true
		case CommandOff:
			state = false
		}

		// Ignore cuboids that are fully out of range
		if step.CoordA[0] > 50 && step.CoordB[0] > 50 {
			continue
		}
		if step.CoordA[1] > 50 && step.CoordB[1] > 50 {
			continue
		}
		if step.CoordA[2] > 50 && step.CoordB[2] > 50 {
			continue
		}
		if step.CoordA[0] < -50 && step.CoordB[0] < -50 {
			continue
		}
		if step.CoordA[1] < -50 && step.CoordB[1] < -50 {
			continue
		}
		if step.CoordA[2] < -50 && step.CoordB[2] < -50 {
			continue
		}

		// Clamp remaining cubeoid to be within range to save on compute
		step.CoordA[0] = Clamp(step.CoordA[0], -50, 50)
		step.CoordA[1] = Clamp(step.CoordA[1], -50, 50)
		step.CoordA[2] = Clamp(step.CoordA[2], -50, 50)
		step.CoordB[0] = Clamp(step.CoordB[0], -50, 50)
		step.CoordB[1] = Clamp(step.CoordB[1], -50, 50)
		step.CoordB[2] = Clamp(step.CoordB[2], -50, 50)

		for x := step.CoordA[0]; x <= step.CoordB[0]; x++ {
			for y := step.CoordA[1]; y <= step.CoordB[1]; y++ {
				for z := step.CoordA[2]; z <= step.CoordB[2]; z++ {
					reactor[Coord3D{x, y, z}] = state
				}
			}
		}
	}

	count := 0

	for _, state := range reactor {
		if state {
			count += 1
		}
	}

	return count
}

func ParseInput(input string) ([]RebootStep, error) {
	steps := []RebootStep{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")

		var command Command
		switch parts[0] {
		case "on":
			command = CommandOn
		case "off":
			command = CommandOff
		}

		// x=10..12,y=10..12,z=10..12
		coordRanges := strings.Split(parts[1], ",")

		// ["x=10..12", "y=10..12", "z=10..12"]
		coordRanges[0] = strings.ReplaceAll(coordRanges[0], "x=", "")
		coordRanges[1] = strings.ReplaceAll(coordRanges[1], "y=", "")
		coordRanges[2] = strings.ReplaceAll(coordRanges[2], "z=", "")

		// ["10..12", "10..12", "10..12"]
		xCoords := strings.Split(coordRanges[0], "..")
		yCoords := strings.Split(coordRanges[1], "..")
		zCoords := strings.Split(coordRanges[2], "..")

		// [["10", "12"], ["10", "12"], ["10", "12"]]
		x1, err := strconv.Atoi(xCoords[0])
		if err != nil {
			return []RebootStep{}, err
		}
		x2, err := strconv.Atoi(xCoords[1])
		if err != nil {
			return []RebootStep{}, err
		}
		y1, err := strconv.Atoi(yCoords[0])
		if err != nil {
			return []RebootStep{}, err
		}
		y2, err := strconv.Atoi(yCoords[1])
		if err != nil {
			return []RebootStep{}, err
		}
		z1, err := strconv.Atoi(zCoords[0])
		if err != nil {
			return []RebootStep{}, err
		}
		z2, err := strconv.Atoi(zCoords[1])
		if err != nil {
			return []RebootStep{}, err
		}

		steps = append(steps, RebootStep{
			Command: command,
			CoordA: Coord3D{
				x1, y1, z1,
			},
			CoordB: Coord3D{
				x2, y2, z2,
			},
		})
	}

	return steps, nil
}

func Clamp(number, lower, upper int) int {

	if number > upper {
		number = upper
	} else if number < lower {
		number = lower
	}

	return number
}
