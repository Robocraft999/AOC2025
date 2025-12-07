package main

import (
	"fmt"
	"slices"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(7)
	input = helper.TestInput(7, "testinput.txt")
	//input = helper.TestInput(7, "testinput2.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	grid := helper.MakeGridFromLines(lines)
	helper.PrintGrid(grid)

	startIndex := strings.Index(lines[0], "S")
	currentIndices := []int{startIndex}
	nextIndices := make([]int, 0)
	sum := 0
	for y := range grid[1:] {
		nextIndices = make([]int, 0)
		for x := range grid[y] {
			if slices.Contains(currentIndices, x) {
				if grid[y][x] == "^" {
					sum += 1
					if x > 0 && !slices.Contains(nextIndices, x-1) {
						nextIndices = append(nextIndices, x-1)
					}
					if x < len(grid[y])-1 && !slices.Contains(nextIndices, x+1) {
						nextIndices = append(nextIndices, x+1)
					}
				} else {
					nextIndices = append(nextIndices, x)
				}
			}
		}
		fmt.Println(nextIndices)
		currentIndices = nextIndices
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	grid := helper.MakeGridFromLines(lines)

	startIndex := strings.Index(lines[0], "S")
	currentIndices := []int{startIndex}
	nextIndices := make([]int, 0)
	sum := 0
	currentIndicesMap := make(map[int]int)
	currentIndicesMap[startIndex] = 1
	for y := range grid[1 : len(grid)-1] {
		y += 1
		nextIndices = make([]int, 0)
		fmt.Println(grid[y], currentIndicesMap)
		for _, x := range currentIndices {
			if slices.Contains(currentIndices, x) {
				if grid[y][x] == "^" {
					amt := currentIndicesMap[x]
					currentIndicesMap[x] = 0
					if x > 0 {
						if !slices.Contains(nextIndices, x-1) {
							nextIndices = append(nextIndices, x-1)
						}
						currentIndicesMap[x-1] += amt
					}
					if x < len(grid[y])-1 {
						if !slices.Contains(nextIndices, x+1) {
							nextIndices = append(nextIndices, x+1)
						}
						currentIndicesMap[x+1] += amt
					}
				} else {
					nextIndices = append(nextIndices, x)
				}
			}
		}
		currentIndices = nextIndices
	}
	fmt.Println(grid[len(grid)-1], currentIndicesMap)
	for _, v := range currentIndicesMap {
		sum += v
	}
	fmt.Println(sum)
}
