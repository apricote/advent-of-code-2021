package main

import (
	"strconv"
	"strings"
)

const (
	DirectionUp      = "up"
	DirectionDown    = "down"
	DirectionForward = "forward"
)

type Command struct {
	Direction string
	Value     int
}

func GetDivePosition(input string) int {
	commands := ParseDay2Input(input)

	var depth int
	var horizontal int

	// Simulate commands
	for _, command := range commands {
		switch command.Direction {
		case DirectionUp:
			depth -= command.Value
		case DirectionDown:
			depth += command.Value
		case DirectionForward:
			horizontal += command.Value
		}
	}

	return depth * horizontal
}

func ParseDay2Input(input string) []Command {
	// Split input into lines
	commandsRaw := strings.Split(input, "\n")

	// Parse commands to datastructure
	commands := make([]Command, 0, len(commandsRaw))
	for _, command := range commandsRaw {
		commandParts := strings.Split(command, " ")

		direction := commandParts[0]
		value, _ := strconv.Atoi(commandParts[1])

		commands = append(commands, Command{
			Direction: direction,
			Value:     value,
		})
	}

	return commands
}
