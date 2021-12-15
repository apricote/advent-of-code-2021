package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetCostToAvoidChitons(input string) int {
	riskLevels, err := ParseInput(input)

	if err != nil {
		log.Fatal(err)
	}

	if len(riskLevels) == 0 {
		log.Fatalf("no entries in riskLevels: %v", riskLevels)
	}

	source := [2]int{0, 0}
	destination := [2]int{len(riskLevels[0]) - 1, len(riskLevels) - 1}

	return Dijkstra(riskLevels, source, destination)
}

func ParseInput(input string) ([][]int, error) {
	riskLevels := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		levelsInLine := []int{}

		for _, levelString := range strings.Split(line, "") {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				return [][]int{}, fmt.Errorf("unable to parse risk level %v", levelString)
			}

			levelsInLine = append(levelsInLine, level)
		}

		riskLevels = append(riskLevels, levelsInLine)
	}

	return riskLevels, nil
}

func Dijkstra(riskLevels [][]int, source [2]int, destination [2]int) int {
	// create vertex set Q
	Q := make(PriorityQueue, 0)

	dist := map[[2]int]int{}

	maxX := len(riskLevels[0]) - 1
	maxY := len(riskLevels) - 1

	// for each vertex v in Graph:
	for y, row := range riskLevels {
		for x := range row {
			v := [2]int{x, y}
			// dist[v] ← INFINITY
			dist[v] = math.MaxInt

			// dist[source] ← 0
			if v == source {
				dist[v] = 0
			}

			// add v to Q
			heap.Push(&Q, &Item{
				value:    v,
				priority: dist[v],
				index:    -1,
			})
		}
	}

	// while Q is not empty:
	for Q.Len() > 0 {
		// u ← vertex in Q with min dist[u]
		// remove u from Q
		u := heap.Pop(&Q).(*Item)

		// Exit early if the current node is equal to destination node
		if u.value == destination {
			continue
		}

		// for each neighbor v of u still in Q:
		for _, v := range GetPossibleNeighbors(u.value, maxX, maxY) {
			// alt ← dist[u] + length(u, v)
			alt := dist[u.value] + riskLevels[v[0]][v[1]]

			// if alt < dist[v]:
			if alt < dist[v] {
				// dist[v] ← alt
				dist[v] = alt

				if vItem := Q.Find(v); vItem != nil {
					Q.update(vItem, v, dist[v])
				}

			}
		}
	}

	return dist[destination]
}

func GetPossibleNeighbors(coords [2]int, maxX, maxY int) [][2]int {
	neighborOffsets := [][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	neighbors := [][2]int{}

	for _, offset := range neighborOffsets {
		newX := coords[0] + offset[0]
		newY := coords[1] + offset[1]

		if newX < 0 || newX > maxX || newY < 0 || newY > maxY {
			continue
		}

		neighbors = append(neighbors, [2]int{newX, newY})
	}

	return neighbors
}
