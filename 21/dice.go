package main

type Dice interface {
	Roll() int
}

type DeterministicDice struct {
	Rolls int
}

func (dd *DeterministicDice) Roll() int {
	dd.Rolls += 1

	return dd.Rolls % 100
}
