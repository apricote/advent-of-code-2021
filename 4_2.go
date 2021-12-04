package main

func GetFinalBingoScoreLosingBoard(rawInput string) int {
	input := ParseDay4Input(rawInput)

	ongoingBoards := make([]BingoBoard, len(input.Boards))
	copy(ongoingBoards, input.Boards)

	var winningBoard BingoBoard
	var winningNumber int

out:
	for i := 0; i < len(input.DrawnNumbers); i++ {
		newOngoingBoards := make([]BingoBoard, 0, len(ongoingBoards))

		for _, board := range ongoingBoards {

			MarkBingoBoard(&board, input.DrawnNumbers[i])

			if CheckBingoWinCondition(board) {
				if len(ongoingBoards) == 1 {
					winningBoard = board
					winningNumber = input.DrawnNumbers[i]
					break out
				}
			} else {
				newOngoingBoards = append(newOngoingBoards, board)
			}

		}

		ongoingBoards = newOngoingBoards
	}

	return CalculateFinalBingoScore(winningBoard, winningNumber)
}
