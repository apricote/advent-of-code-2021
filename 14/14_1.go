package main

import (
	"math"
	"strings"
)

type Input struct {
	Template []rune
	Pairs    Pairs
	Rules    Rules
}

type Rules map[string][2]string
type Pairs map[string]int

func GetCommonElementCount(inputString string) int {
	input := ParseInput(inputString)

	for i := 0; i < 10; i++ {
		input.Pairs = ExpandTemplate(input.Pairs, input.Rules)
	}

	return GetLeastMostValue(input.Template, input.Pairs)
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

		insert := string(ruleSections[1][0])

		rules[ruleSections[0]] = [2]string{
			string(ruleSections[0][0]) + insert,
			insert + string(ruleSections[0][1]),
		}
	}

	// Pairs
	pairs := Pairs{}
	for i, cur := range template {
		if i == len(template)-1 {
			break
		}

		next := template[i+1]

		pairs[string(cur)+string(next)] += 1
	}

	return Input{
		Template: template,
		Rules:    rules,
		Pairs:    pairs,
	}
}

func ExpandTemplate(pairs Pairs, rules Rules) Pairs {
	newPairs := Pairs{}

	for pair, count := range pairs {

		resultingPairs, ok := rules[pair]

		if !ok {
			newPairs[pair] += count
		} else {
			for _, newPair := range resultingPairs {
				newPairs[newPair] += count
			}
		}
	}

	return newPairs
}

func CountElement(template []rune, pairs Pairs) map[rune]int {
	countsFromPairs := map[rune]int{}
	for pair, count := range pairs {
		countsFromPairs[rune(pair[0])] += count
		countsFromPairs[rune(pair[1])] += count
	}

	firstElement := template[0]
	lastElement := template[len(template)-1]

	counts := map[rune]int{}
	for element, count := range countsFromPairs {
		if element == firstElement {
			count -= 1
			counts[element] += 1
		}

		if element == lastElement {
			count -= 1
			counts[element] += 1
		}

		counts[element] += count / 2
	}

	return counts
}

func GetLeastMostValue(template []rune, pairs Pairs) int {
	counts := CountElement(template, pairs)

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
