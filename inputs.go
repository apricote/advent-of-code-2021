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
