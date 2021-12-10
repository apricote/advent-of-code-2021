package main

import (
	"strings"
)

type Stack []string

func GetSyntaxErrorScore(input string) int {
	lines := ParseInput(input)

	points := 0

	for _, line := range lines {
		_, illegalCharacter := ParseLine(line)
		points += GetIllegalCharacterPoints(illegalCharacter)
	}

	return points
}

func ParseInput(input string) [][]string {
	parsedLines := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		parsedLines = append(parsedLines, append([]string{}, strings.Split(line, "")...))
	}

	return parsedLines
}

// Returns the first illegal character or "" in case the line is incomplete
func ParseLine(line []string) (Stack, string) {
	stack := Stack{}

	for _, token := range line {
		lastElement := ""

		switch token {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			fallthrough
		case "<":
			stack = append(stack, token)
		case ")":
			stack, lastElement = stack.Pop()

			if lastElement != "(" {
				return stack, token
			}
		case "]":
			stack, lastElement = stack.Pop()

			if lastElement != "[" {
				return stack, token
			}
		case "}":
			stack, lastElement = stack.Pop()

			if lastElement != "{" {
				return stack, token
			}
		case ">":
			stack, lastElement = stack.Pop()

			if lastElement != "<" {
				return stack, token
			}
		}
	}

	return stack, ""
}

func GetIllegalCharacterPoints(illegalChar string) int {
	switch illegalChar {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func (s Stack) Pop() (Stack, string) {
	topIndex := len(s) - 1

	value := s[topIndex]
	s[topIndex] = ""
	s = s[:topIndex]

	return s, value
}
