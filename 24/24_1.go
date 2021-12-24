package main

import (
	"log"

	"github.com/apricote/advent-of-code-2021/util"
)

var (
	monadInstructions = ParseProgram(util.GetInput())
)

func MONAD(modelNumber int) bool {
	inputs := []int{}
	for modelNumber > 0 {
		inputs = append(inputs, modelNumber%10)
		modelNumber /= 10
	}

	for i := 0; i < len(inputs)/2; i++ {
		inputs[i], inputs[len(inputs)-1-i] = inputs[len(inputs)-1-i], inputs[i]
	}

	registers := RunALU(&monadInstructions, inputs)
	log.Printf("Registers: %v", registers)

	return registers[3] == 0
}

func BruteForceModelNumber() int {
	for i := 99999999999999; i >= 11111111111111; i-- {
		if i%10000 == 0 {
			log.Printf("Checked %v (%v)", i, (float64(99999999999999-i) / 88888888888888))
		}

		if MONAD(i) {
			return i
		}
	}

	return 0
}
