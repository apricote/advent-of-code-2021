package main

import (
	"fmt"
	"log"
)

// 1,1,1 => 3
// 1,1,2 => 4
// 1,1,3 => 5

// 1,2,1 => 4
// 1,2,2 => 5
// 1,2,3 => 6

// 1,3,1 => 5
// 1,3,2 => 6
// 1,3,3 => 7

// 2,1,1 => 4
// 2,1,2 => 5
// 2,1,3 => 6

// 2,2,1 => 5
// 2,2,2 => 6
// 2,2,3 => 7

// 2,3,1 => 6
// 2,3,2 => 7
// 2,3,3 => 8

// 3,1,1 => 5
// 3,1,2 => 6
// 3,1,3 => 7

// 3,2,1 => 6
// 3,2,2 => 7
// 3,2,3 => 8

// 3,3,1 => 7
// 3,3,2 => 8
// 3,3,3 => 9

// 3 => 1
// 4 => 3
// 5 => 6
// 6 => 7
// 7 => 6
// 8 => 3
// 9 => 1

var (
	quantumDiceResults = map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
)

type GameState map[[4]int]int

func GetDiracQuantumDiceGameScore(input string) int {
	startingPositions, err := ParseInput(input)
	if err != nil {
		log.Fatalf("unable to parse input: %v", err)
	}

	gameState := QuantumDiracDiceGame(startingPositions)

	universesOfWinner, err := CountUniversesOfWinner(gameState)
	if err != nil {
		log.Fatalf("unable to determine universe count for winner: %v", err)
	}

	return universesOfWinner
}

func QuantumDiracDiceGame(startingPositions [2]int) GameState {
	gameState := GameState{
		[4]int{startingPositions[0], 0, startingPositions[1], 0}: 1,
	}

	for {
		// Do turns until no universe is undecided
		allDecided := true
		for state := range gameState {
			if GetWinner(state) == -1 {
				allDecided = false
				break
			}
		}
		if allDecided {
			break
		}

		gameState = QuantumTurn(gameState, 1)
		gameState = QuantumTurn(gameState, 2)
	}

	return gameState
}

func QuantumTurn(gameState GameState, playerIndex int) GameState {
	newGameState := GameState{}

	for state, count := range gameState {
		if GetWinner(state) > 0 {
			newGameState[state] += count
			continue
		}
		// If state is decided => continue

		for rolls, universes := range quantumDiceResults {
			newState := state

			playerPositionIndex := -1
			playerScoreIndex := -1

			if playerIndex == 1 {
				playerPositionIndex = 0
				playerScoreIndex = 1
			} else {
				playerPositionIndex = 2
				playerScoreIndex = 3
			}

			// Update Position
			newState[playerPositionIndex] += rolls
			newState[playerPositionIndex] = ((newState[playerPositionIndex] - 1) % 10) + 1

			// Update score
			newState[playerScoreIndex] += newState[playerPositionIndex]

			newGameState[newState] += universes * count
		}
	}

	return newGameState
}

// -1 => No Winner
// 1  => Player1
// 2  => Player2
func GetWinner(state [4]int) int {
	if state[1] >= 21 {
		return 1
	}

	if state[3] >= 21 {
		return 2
	}

	return -1
}

func CountUniversesOfWinner(gameState GameState) (int, error) {
	counts := map[int]int{
		1: 0, // Universes where Player 1 won
		2: 0, // Universes where Player 2 won
	}

	for state, universeCount := range gameState {
		winner := GetWinner(state)
		if winner == -1 {
			return 0, fmt.Errorf("undecided universes")
		}

		counts[winner] += universeCount

	}

	max := 0

	if counts[1] > counts[2] {
		max = counts[1]
	} else {
		max = counts[2]
	}

	return max, nil
}
