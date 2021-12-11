package main

func GetFirstDayWithSynchronizedFlashes(input string) int {
	energyMap := ParseInput(input)

	step := 0
	flashesInPreviousStep := 0

	for flashesInPreviousStep != EMSize*EMSize {
		step += 1
		flashesInPreviousStep = energyMap.SimulateStep()
	}

	return step
}
