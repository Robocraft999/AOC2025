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
	//input = helper.TestInput(10, "testinput.txt")
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

func part2(input string) {
	fmt.Println("Part 2:", "See python")
}
