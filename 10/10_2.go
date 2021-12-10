package main

import "sort"

func GetCompletionScore(input string) int {
	lines := ParseInput(input)

	completionScores := []int{}

	for _, line := range lines {
		stack, illegalCharacter := ParseLine(line)

		if illegalCharacter == "" {
			completionScores = append(completionScores, stack.CompletionScore())
		}
	}

	sort.Ints(completionScores)

	return completionScores[len(completionScores)/2]
}

func (s Stack) CompletionScore() int {
	score := 0

	for len(s) > 0 {
		token := ""
		s, token = s.Pop()

		switch token {
		case "(":
			score *= 5
			score += 1
		case "[":
			score *= 5
			score += 2

		case "{":
			score *= 5
			score += 3

		case "<":
			score *= 5
			score += 4
		}
	}

	return score
}
