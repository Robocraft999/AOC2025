package main

import (
	"cmp"
	"fmt"
	"slices"
	sort2 "sort"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(8)
	//input = helper.TestInput(8, "testinput.txt")
	part1(input)
	part2(input)
}

type DistPair struct {
	start [3]int
	end   [3]int
	dist  float64
}

func sort(first DistPair, second DistPair) int {
	return cmp.Compare(first.dist, second.dist)
}

func addToExistingOrNewCircuit(circuits *[][][3]int, pair DistPair) bool {
	startCircuitIndex := -1
	endCircuitIndex := -1
	for i, circuit := range *circuits {
		startIn := slices.Contains(circuit, pair.start)
		endIn := slices.Contains(circuit, pair.end)
		if startIn && endIn {
			return false
		} else if startIn {
			startCircuitIndex = i
		} else if endIn {
			endCircuitIndex = i
		}
	}
	if startCircuitIndex != -1 && endCircuitIndex != -1 {
		firstIndex := startCircuitIndex
		secondIndex := endCircuitIndex
		if startCircuitIndex > endCircuitIndex {
			firstIndex = endCircuitIndex
			secondIndex = startCircuitIndex
		}
		for _, coord := range (*circuits)[secondIndex] {
			(*circuits)[firstIndex] = append((*circuits)[firstIndex], coord)
		}
		(*circuits)[secondIndex] = [][3]int{}
		return true
	} else if startCircuitIndex != -1 {
		(*circuits)[startCircuitIndex] = append((*circuits)[startCircuitIndex], pair.end)
		return true
	} else if endCircuitIndex != -1 {
		(*circuits)[endCircuitIndex] = append((*circuits)[endCircuitIndex], pair.start)
		return true
	}
	*circuits = append(*circuits, [][3]int{pair.start, pair.end})
	return true
}

func part1(input string) {
	fmt.Println("Part 1:")
	lines := strings.Split(input, "\n")
	coords := make([][3]int, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		z, _ := strconv.Atoi(splitLine[2])
		coords[i] = [3]int{x, y, z}
	}
	distPairs := make([]DistPair, 0)
	maxPairs := 1000
	//maxPairs := 10
	for _, first := range coords {
		for _, coord := range coords {
			distance := helper.EuclideanDistance(first, coord)
			if distance > 0 {
				if !slices.Contains(distPairs, DistPair{coord, first, distance}) {
					if len(distPairs) < maxPairs {
						distPairs = append(distPairs, DistPair{first, coord, distance})
						slices.SortFunc(distPairs, sort)
					} else {
						for _, distPair := range distPairs {
							if distPair.dist > distance {
								distPairs = append(distPairs, DistPair{first, coord, distance})
								slices.SortFunc(distPairs, sort)
								distPairs = distPairs[:maxPairs]
								break
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(distPairs)
	circuits := make([][][3]int, 0)
	for _, pair := range distPairs {
		addToExistingOrNewCircuit(&circuits, pair)
	}
	sizeComp := func(first [][3]int, second [][3]int) int {
		return cmp.Compare(len(second), len(first))
	}
	slices.SortFunc(circuits, sizeComp)
	fmt.Println(len(circuits[0]))
	fmt.Println(len(circuits[1]))
	fmt.Println(len(circuits[2]))
	fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}

func part2(input string) {
	fmt.Println("Part 2:")
	lines := strings.Split(input, "\n")
	coords := make([][3]int, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		z, _ := strconv.Atoi(splitLine[2])
		coords[i] = [3]int{x, y, z}
	}
	distPairMap := make(map[[2][3]int]float64)
	for _, first := range coords {
		for _, coord := range coords {
			distance := helper.EuclideanDistance(first, coord)
			if distance > 0 {
				if _, ok := distPairMap[[2][3]int{coord, first}]; !ok {
					distPairMap[[2][3]int{coord, first}] = distance
				}
			}
		}
	}
	keys := make([][2][3]int, 0, len(distPairMap))

	for key := range distPairMap {
		keys = append(keys, key)
	}
	sort2.SliceStable(keys, func(i, j int) bool {
		return distPairMap[keys[i]] < distPairMap[keys[j]]
	})
	circuits := make([][][3]int, 0)
	sizeComp := func(first [][3]int, second [][3]int) int {
		return cmp.Compare(len(second), len(first))
	}
	for _, key := range keys {
		start := key[0]
		end := key[1]
		addToExistingOrNewCircuit(&circuits, DistPair{start, end, distPairMap[key]})
		slices.SortFunc(circuits, sizeComp)
		if len(circuits[0]) == len(coords) {
			fmt.Println(start, end)
			fmt.Println(start[0] * end[0])
			break
		}
	}
}
