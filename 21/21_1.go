package main

import (
	"log"
	"strconv"
	"strings"
)

func GetDiracDiceGameScore(input string) int {
	startingPositions, err := ParseInput(input)
	if err != nil {
		log.Fatalf("Unable to parse dirac dice game input: %v", err)
	}

	scores := DiracDiceGame(startingPositions)

	lowerScore := 0

	if scores[0] < scores[1] {
		lowerScore = scores[0]
	} else {
		lowerScore = scores[1]
	}

	return lowerScore * scores[2]
}

func ParseInput(input string) ([2]int, error) {
	positionStrings := strings.Split(input, "\n")

	player1Position, err := strconv.Atoi(string(positionStrings[0][28]))
	if err != nil {
		return [2]int{}, err
	}

	player2Position, err := strconv.Atoi(string(positionStrings[1][28]))
	if err != nil {
		return [2]int{}, err
	}

	return [2]int{
		player1Position,
		player2Position,
	}, nil
}

// DiracDiceGame returns an array with 3 values:
// 0 -> Player 1 Score
// 1 -> Player 2 Score
// 2 -> Dice Rolls
func DiracDiceGame(startingPositions [2]int) [3]int {
	player1Score := 0
	player2Score := 0

	player1Position := startingPositions[0]
	player2Position := startingPositions[1]

	dice := DeterministicDice{}

	for {
		player1Position = Turn(player1Position, &dice)
		player1Score += player1Position

		if player1Score >= 1000 {
			// Player 1 has Won!
			break
		}

		player2Position = Turn(player2Position, &dice)
		player2Score += player2Position

		if player2Score >= 1000 {
			// Player 2 has Won!
			break
		}
	}

	return [3]int{
		player1Score,
		player2Score,
		dice.Rolls,
	}
}

func Turn(position int, dice Dice) int {
	rolls := dice.Roll() + dice.Roll() + dice.Roll()

	position += rolls
	// available positions are in interval (1,10)
	position = ((position - 1) % 10) + 1

	return position
}
