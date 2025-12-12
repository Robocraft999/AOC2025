package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(12)
	//too small to work
	//input = helper.TestInput(12, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	partsLen := len(parts)
	shapesS := parts[:partsLen-1]
	regions := strings.Split(parts[partsLen-1], "\n")
	shapes := make([][]string, len(shapesS))
	shapesSpace := make([]int, len(shapesS))
	for i, shape := range shapesS {
		shapes[i] = strings.Split(shape, "\n")[1:]
		shapesSpace[i] = strings.Count(shape, "#")
	}
	fmt.Println(shapes, len(shapes))

	sum := 0

	for _, region := range regions {
		regionParts := strings.Split(region, ": ")
		for _, part := range regionParts {
			fmt.Println(part)
		}
		fmt.Println(len(regionParts))
		sizeS := strings.Split(regionParts[0], "x")
		w, _ := strconv.Atoi(sizeS[0])
		h, _ := strconv.Atoi(sizeS[1])
		size := [2]int{w, h}
		amounts := make([]int, 0)
		for _, rp := range strings.Split(regionParts[1], " ") {
			num, _ := strconv.Atoi(rp)
			amounts = append(amounts, num)
		}
		fmt.Println(size, amounts)
		space := 0
		for i, amt := range amounts {
			space += amt * shapesSpace[i]
		}
		if space <= size[0]*size[1] {
			sum += 1
		}
	}

	fmt.Println("Part1", sum)
}

func part2(input string) {
	fmt.Println("Click")
}
