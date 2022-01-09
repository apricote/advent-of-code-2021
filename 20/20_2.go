package main

import "log"

func EnhanceImageFiftyTimes(inputString string) int {
	input := ParseInput(inputString)

	image := input.Image

	for x := 0; x < 50; x++ {
		image = image.Enhance(input.ImageEnhancementAlgorithm)

	}

	litPixels, err := image.CountLitPixels()
	if err != nil {
		log.Fatalf("Unable to count lit pixels: %v", err)
	}

	return litPixels
}
