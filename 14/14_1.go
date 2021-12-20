package main

import (
	"math"
	"strings"
	"unicode/utf8"
)

type Input struct {
	Template []rune
	Rules    Rules
}

type Rules map[string]rune

func GetCommonElementCount(inputString string) int {
	input := ParseInput(inputString)

	for i := 0; i < 10; i++ {
		input.Template = ExpandTemplate(input.Template, input.Rules)
	}

	return GetLeastMostValue(input.Template)
}

func ParseInput(input string) Input {
	sections := strings.Split(input, "\n\n")

	// Template
	template := []rune{}
	for _, c := range sections[0] {
		template = append(template, c)
	}

	// Rules
	rules := Rules{}
	for _, ruleString := range strings.Split(sections[1], "\n") {
		ruleSections := strings.Split(ruleString, " -> ")

		insert, _ := utf8.DecodeRuneInString(ruleSections[1])

		rules[ruleSections[0]] = insert
	}

	return Input{
		Template: template,
		Rules:    rules,
	}
}

func ExpandTemplate(tpl []rune, rules Rules) []rune {
	newTpl := []rune{}

	for i, cur := range tpl {
		newTpl = append(newTpl, cur)

		if i == len(tpl)-1 {
			break
		}

		// Not last rune, let's check upcoming pair
		next := tpl[i+1]

		insertElement, ok := rules[string(cur)+string(next)]

		if ok {
			// matching rule
			newTpl = append(newTpl, insertElement)
		}
	}

	return newTpl
}

func CountElement(tpl []rune) map[rune]int {
	counts := map[rune]int{}

	for _, cur := range tpl {
		counts[cur] += 1
	}

	return counts
}

func GetLeastMostValue(tpl []rune) int {
	counts := CountElement(tpl)

	leastCommon := math.MaxInt
	mostCommon := math.MinInt

	for _, value := range counts {
		if value < leastCommon {
			leastCommon = value
		}

		if value > mostCommon {
			mostCommon = value
		}
	}

	return mostCommon - leastCommon
}
