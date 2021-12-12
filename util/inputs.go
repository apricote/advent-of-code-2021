package util

import (
	"log"
	"os"
)

func GetInputFromPath(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("error while reading %s", path)
	}

	return string(dat)
}

func GetExampleInput() string {
	return GetInputFromPath("input_example.txt")
}

func GetInput() string {
	return GetInputFromPath("input.txt")
}
