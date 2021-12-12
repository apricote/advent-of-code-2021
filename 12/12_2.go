package main

func GetCaveTimesWithAdditionalTime(input string) int {
	caves := ParseInput(input)

	isAcceptableNextCave := func(path CavePath, nextCave *Cave) bool {
		if nextCave.Name == "start" {
			return false
		}

		if nextCave.Big {
			return true
		}

		hasVisitedSmallCaveTwice := false
		for _, cave := range path {
			if cave.Big {
				continue
			}

			if ContainsCaveTimes(path, cave) == 2 {
				hasVisitedSmallCaveTwice = true
				break
			}
		}

		if !hasVisitedSmallCaveTwice && ContainsCaveTimes(path, nextCave) < 2 {
			return true
		}

		if hasVisitedSmallCaveTwice && ContainsCaveTimes(path, nextCave) < 1 {
			return true
		}

		return false
	}

	paths := GetPossiblePaths(caves, isAcceptableNextCave)

	return len(paths)
}
