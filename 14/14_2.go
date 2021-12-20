package main

func GetCommonElementCountAfter40Rounds(inputString string) int {
	input := ParseInput(inputString)

	for i := 0; i < 40; i++ {
		input.Pairs = ExpandTemplate(input.Pairs, input.Rules)
		//log.Printf("Iteration %v: Length is %v", i, len(input.Pairs))
	}

	return GetLeastMostValue(input.Template, input.Pairs)
}
