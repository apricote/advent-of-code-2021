package main

func GetDivePositionWithAim(input string) int {
	commands := ParseDay2Input(input)

	var aim int
	var depth int
	var horizontal int

	for _, command := range commands {
		switch command.Direction {
		case DirectionDown:
			aim += command.Value
		case DirectionUp:
			aim -= command.Value
		case DirectionForward:
			horizontal += command.Value
			depth += aim * command.Value
		}
	}

	return depth * horizontal
}
