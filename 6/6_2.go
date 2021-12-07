package main

func CountLanternfish256(input string) int {
	ages := ParseDay6Input(input)

	for i := 0; i < 256; i++ {
		ages = SimulateLanternfishes(ages)
	}

	return CountAllLanternfish(ages)
}
