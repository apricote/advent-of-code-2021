package util

import (
	"log"
	"os"
)

func getInput(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("error while reading %s", path)
	}

	return string(dat)
}

func GetExampleInput() string {
	return getInput("input_example.txt")
}

func GetInput() string {
	return getInput("input.txt")
}
