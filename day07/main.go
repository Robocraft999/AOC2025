package main

import (
	"fmt"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(7)
	//input = helper.TestInput(7, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	grid := helper.MakeGridFromLines(lines)
	helper.PrintGrid(grid)

	startIndex := strings.Index(lines[0], "S")
	sum := 0
	currentIndicesMap := make(map[int]int)
	currentIndicesMap[startIndex] = 1
	for y := range grid[1 : len(grid)-1] {
		y += 1
		for x, amt := range currentIndicesMap {
			if amt > 0 {
				if grid[y][x] == "^" {
					currentIndicesMap[x] = 0
					sum += 1
					if x > 0 {
						currentIndicesMap[x-1] += amt
					}
					if x < len(grid[y])-1 {
						currentIndicesMap[x+1] += amt
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	grid := helper.MakeGridFromLines(lines)

	startIndex := strings.Index(lines[0], "S")
	sum := 0
	currentIndicesMap := make(map[int]int)
	currentIndicesMap[startIndex] = 1
	for y := range grid[1 : len(grid)-1] {
		y += 1
		for x, amt := range currentIndicesMap {
			if amt > 0 {
				if grid[y][x] == "^" {
					currentIndicesMap[x] = 0
					if x > 0 {
						currentIndicesMap[x-1] += amt
					}
					if x < len(grid[y])-1 {
						currentIndicesMap[x+1] += amt
					}
				}
			}
		}
	}
	for _, v := range currentIndicesMap {
		sum += v
	}
	fmt.Println(sum)
}
