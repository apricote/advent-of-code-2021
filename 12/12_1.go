package main

import (
	"strings"
)

type Cave struct {
	Name        string
	Big         bool
	Connections []*Cave
}

type CaveMap map[string]*Cave

type CavePath []*Cave

type NextCaveDecider func(CavePath, *Cave) bool

func GetCavePaths(input string) int {
	caves := ParseInput(input)

	isAcceptableNextCave := func(path CavePath, nextCave *Cave) bool {
		return !nextCave.Big && ContainsCave(path, nextCave)
	}

	paths := GetPossiblePaths(caves, isAcceptableNextCave)

	return len(paths)
}

func ParseInput(input string) CaveMap {
	caves := CaveMap{}

	for _, line := range strings.Split(input, "\n") {
		connectedCaves := strings.Split(line, "-")

		// Get the Cave struct for connected caves
		cave1 := GetOrCreateCave(&caves, connectedCaves[0])
		cave2 := GetOrCreateCave(&caves, connectedCaves[1])

		// Add both caves to eathothers Connections
		cave1.Connections = append(cave1.Connections, cave2)
		cave2.Connections = append(cave2.Connections, cave1)
	}

	return caves
}

func CaveIsBig(name string) bool {
	return strings.ToUpper(name) == name
}

func GetOrCreateCave(caves *CaveMap, name string) *Cave {
	cave, ok := (*caves)[name]
	if !ok {
		cave = &Cave{
			Name:        name,
			Big:         CaveIsBig(name),
			Connections: []*Cave{},
		}
		(*caves)[cave.Name] = cave
	}

	return cave
}

func GetPossiblePaths(caveMap CaveMap, isAcceptableNextCave NextCaveDecider) []CavePath {
	startCave := caveMap["start"]
	endCave := caveMap["end"]

	endingPaths := []CavePath{}

	// Start exploration from startCave
	viablePaths := []CavePath{{startCave}}

	// Loop this until finished(?)
	for len(viablePaths) > 0 {
		nextViablePaths := []CavePath{}

		for _, viablePath := range viablePaths {
			lastCaveInPath := viablePath[len(viablePath)-1]

			nextCaves := lastCaveInPath.Connections
			for _, nextCave := range nextCaves {
				// If next cave is endCave that path is considered "finished"
				// and moved to endingPaths
				if nextCave.IsEqual(*endCave) {
					endingPaths = append(endingPaths, AddToPath(viablePath, nextCave))
					continue
				}

				// If next cave is small, check that it's not already in path
				if isAcceptableNextCave(viablePath, nextCave) {
					// Skip adding to viablePaths
					continue
				}

				nextViablePaths = append(nextViablePaths, AddToPath(viablePath, nextCave))
			}
		}
		viablePaths = nextViablePaths
	}

	return endingPaths
}

func ContainsCave(caves CavePath, cave *Cave) bool {
	for _, caveB := range caves {
		if cave.IsEqual(*caveB) {
			return true
		}
	}

	return false
}

func (c Cave) IsEqual(b Cave) bool {
	return c.Name == b.Name
}

func (c Cave) String() string {
	return c.Name
}

func AddToPath(path CavePath, next *Cave) CavePath {
	newPath := make(CavePath, len(path)+1)
	copy(newPath, path)

	newPath[len(path)] = next

	return newPath
}
