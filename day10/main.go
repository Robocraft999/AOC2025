package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(10)
	input = helper.TestInput(10, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		patternS := strings.Replace(strings.Replace(parts[0][1:len(parts[0])-1], ".", "0", -1), "#", "1", -1)
		patternChars := strings.Split(patternS, "")
		slices.Reverse(patternChars)
		patternS = strings.Join(patternChars, "")
		patternP, _ := strconv.ParseInt(patternS, 2, 64)
		pattern := int(patternP)

		sequences := make([]int, len(parts)-1)
		for i, seq := range parts[1 : len(parts)-1] {
			num := 0
			for _, d := range strings.Split(seq[1:len(seq)-1], ",") {
				numP, _ := strconv.Atoi(d)
				num += 1 << numP
			}
			sequences[i] = num
		}

		numCombs := helper.PowInt(2, len(sequences))
		minTries := math.MaxInt
		for i := 0; i < numCombs; i++ {
			num := 0
			tries := 0
			for j, seq := range sequences {
				if i&(1<<j) != 0 {
					num ^= seq
					tries += 1
				}
			}
			if num == pattern {
				if tries < minTries {
					minTries = tries
				}
			}
		}
		sum += minTries
	}
	fmt.Println("Part 1:", sum)
}

// greater, equal
func compareSlices(target, current []int) (bool, bool) {
	if len(target) != len(current) {
		panic("not equal length")
	}
	flag := true
	for i := 0; i < len(target); i++ {
		if target[i] > current[i] {
			flag = false
		} else if target[i] < current[i] {
			return true, false
		}
	}
	return false, flag
}

func addOneAtIndices(current, indices []int) []int {
	result := slices.Clone(current)
	for _, i := range indices {
		result[i] += 1
	}
	return result
}

func build(target, current []int, sequences [][]int, depth int) ([]int, int, bool) {
	var result []int
	resDepth := math.MaxInt
	for _, seq := range sequences {
		added := addOneAtIndices(current, seq)
		ge, eq := compareSlices(target, added)
		if ge {
			return added, depth + 1, false
		}
		if eq {
			return added, depth + 1, true
		}
		res, dep, found := build(target, added, sequences, depth+1)
		if found {
			if dep < resDepth {
				resDepth = dep
				result = res
			}
		}
	}
	if result != nil {
		return result, resDepth, true
	}
	return current, depth, false
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		fmt.Println()

		sequences := make([][]int, len(parts)-2)
		for i, seq := range parts[1 : len(parts)-1] {
			numS := strings.Split(seq[1:len(seq)-1], ",")

			indices := make([]int, len(numS))
			for j, d := range numS {
				num, _ := strconv.Atoi(d)
				indices[j] = num
			}
			sequences[i] = indices
		}
		fmt.Println(sequences)

		jols := parts[len(parts)-1]
		jols = jols[1 : len(jols)-1]
		joltagesS := strings.Split(jols, ",")
		joltages := make([]int, len(joltagesS))
		current := make([]int, len(joltagesS))
		for i, joltage := range joltagesS {
			num, _ := strconv.Atoi(joltage)
			joltages[i] = num
			current[i] = 0
		}
		fmt.Println(joltages)
		//res, depth, ok := build(joltages, current, sequences, 0)
		//fmt.Println(res, depth, ok)
	}
	fmt.Println("Part 2:", sum)
}
