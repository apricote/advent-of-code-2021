package main

import (
	"log"
	"strconv"
	"strings"
)

type BingoField struct {
	Number int
	Marked bool
}

type BingoBoard [5][5]*BingoField

type Day4Input struct {
	DrawnNumbers []int
	Boards       []BingoBoard
}

func GetFinalBingoScore(rawInput string) int {
	input := ParseDay4Input(rawInput)

	var winningBoard BingoBoard
	var winningNumber int
out:
	for i := 0; i < len(input.DrawnNumbers); i++ {
		for _, board := range input.Boards {
			MarkBingoBoard(&board, input.DrawnNumbers[i])

			if CheckBingoWinCondition(board) {
				winningBoard = board
				winningNumber = input.DrawnNumbers[i]
				break out
			}
		}
	}

	return CalculateFinalBingoScore(winningBoard, winningNumber)
}

func ParseDay4Input(input string) Day4Input {
	inputLines := strings.Split(input, "\n")

	drawnNumbers := make([]int, 0)
	for _, numberString := range strings.Split(inputLines[0], ",") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Panicf("unable to parse drawn number %v", numberString)
		}

		drawnNumbers = append(drawnNumbers, number)
	}

	numBoards := (len(inputLines) - 1) / 6
	boards := make([]BingoBoard, 0, numBoards)

	for boardNumber := 0; boardNumber < numBoards; boardNumber++ {
		boardOffset := 2 + boardNumber*6
		boardLines := inputLines[boardOffset : boardOffset+5]

		board := BingoBoard{}

		for i, line := range boardLines {
			for j := 0; j < 5; j++ {
				digitOffset := 0 + j*3
				digits := strings.TrimSpace(line[digitOffset : digitOffset+2])

				number, err := strconv.Atoi(digits)
				if err != nil {
					log.Fatalf("unable to parse board number %v", digits)
				}

				board[i][j] = &BingoField{
					Number: number,
					Marked: false,
				}
			}
		}

		boards = append(boards, board)
	}

	return Day4Input{
		DrawnNumbers: drawnNumbers,
		Boards:       boards,
	}
}

func MarkBingoBoard(board *BingoBoard, number int) {
	for _, line := range board {
		for _, field := range line {
			if field.Number == number {
				field.Marked = true
			}
		}
	}
}

func CheckBingoWinCondition(board BingoBoard) bool {
	// Check rows
	for i := 0; i < 5; i++ {
		foundFalse := false
		for _, field := range board[i] {
			if !field.Marked {
				foundFalse = true
				break
			}
		}

		if !foundFalse {
			return true
		}
	}

	// Check columns
	for i := 0; i < 5; i++ {
		foundFalse := false

		for j := 0; j < 5; j++ {
			if !board[j][i].Marked {
				foundFalse = true
				break
			}
		}

		if !foundFalse {
			return true
		}
	}

	return false
}

func CalculateFinalBingoScore(winningBoard BingoBoard, winningNumber int) int {
	sum := 0

	for _, line := range winningBoard {
		for _, field := range line {
			if !field.Marked {
				sum += field.Number
			}
		}
	}

	return sum * winningNumber
}
