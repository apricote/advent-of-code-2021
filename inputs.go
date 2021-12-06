package main

import (
	"log"
	"os"
)

func GetInput(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("error while reading %s", path)
	}

	return string(dat)
}

func GetDay1ExampleInput() string {
	return GetInput("input_1_example.txt")
}

func GetDay1Input() string {
	return GetInput("input_1.txt")
}

func GetDay2ExampleInput() string {
	return GetInput("input_2_example.txt")
}

func GetDay2Input() string {
	return GetInput("input_2.txt")
}

func GetDay3ExampleInput() string {
	return GetInput("input_3_example.txt")
}

func GetDay3Input() string {
	return GetInput("input_3.txt")
}

func GetDay4ExampleInput() string {
	return GetInput("input_4_example.txt")
}

func GetDay4Input() string {
	return GetInput("input_4.txt")
}

func GetDay5ExampleInput() string {
	return GetInput("input_5_example.txt")
}

func GetDay5Input() string {
	return GetInput("input_5.txt")
}

func GetDay6ExampleInput() string {
	return GetInput("input_6_example.txt")
}

func GetDay6Input() string {
	return GetInput("input_6.txt")
}
