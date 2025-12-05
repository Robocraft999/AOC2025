package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(5)
	//input = helper.TestInput(5, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	rangeStrings := strings.Split(parts[0], "\n")
	ranges := make([][]int, len(rangeStrings))
	for i, rangeString := range rangeStrings {
		if rangeString == "" {
			continue
		}
		split := strings.Split(rangeString, "-")
		arr := make([]int, len(split))
		arr[0], _ = strconv.Atoi(split[0])
		arr[1], _ = strconv.Atoi(split[1])
		ranges[i] = arr
	}
	ids := strings.Split(parts[1], "\n")

	fresh := 0
	for _, i := range ids {
		id, _ := strconv.Atoi(i)
		for _, rang := range ranges {
			if id >= rang[0] && id <= rang[1] {
				fresh += 1
				break
			}
		}
	}
	fmt.Println(fresh)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	rangeStrings := strings.Split(parts[0], "\n")
	ranges := make([][]int, len(rangeStrings))
	for i, rangeString := range rangeStrings {
		if rangeString == "" {
			continue
		}
		split := strings.Split(rangeString, "-")
		arr := make([]int, len(split))
		arr[0], _ = strconv.Atoi(split[0])
		arr[1], _ = strconv.Atoi(split[1])
		ranges[i] = arr
	}

	fresh := 0
	calculatedRanges := make([][]int, 1)
	calculatedRanges[0] = []int{0, -1}
	for _, rang := range ranges {
		for _, calculatedRange := range calculatedRanges {
			if rang[0] >= calculatedRange[0] && rang[0] <= calculatedRange[1] && rang[1] > calculatedRange[1] {
				rang[0] = calculatedRange[1] + 1
			} else if rang[1] >= calculatedRange[0] && rang[1] <= calculatedRange[1] && rang[0] < calculatedRange[0] {
				rang[1] = calculatedRange[0] - 1
			} else if rang[0] <= calculatedRange[0] && rang[1] >= calculatedRange[1] {
				calculatedRange[0] = 0
				calculatedRange[1] = -1
			} else if rang[0] >= calculatedRange[0] && rang[1] <= calculatedRange[1] {
				rang[0] = 0
				rang[1] = -1
			}
		}
		calculatedRanges = append(calculatedRanges, rang)
	}
	for _, calculatedRange := range calculatedRanges {
		fresh += calculatedRange[1] - calculatedRange[0] + 1
	}
	fmt.Println(fresh)
}
