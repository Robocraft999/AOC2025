package main

import (
	"fmt"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(4)
	//input = helper.TestInput(4, "testinput.txt")
	part1(input)
	part2(input)
}

func checkPos(grid [][]string, x, y int) bool {
	return helper.CountNeighbours8(grid, x, y) < 4
}

func part1(input string) {
	lines := strings.Fields(input)
	grid := helper.MakeGridFromLines(lines)

	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "@" && checkPos(grid, x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(input string) {
	lines := strings.Fields(input)
	grid := helper.MakeGridFromLines(lines)

	count := 0

	removed := 1
	for removed > 0 {
		removed = 0
		newGrid := make([][]string, len(lines))
		for i := range lines {
			newGrid[i] = make([]string, len(grid[i]))
		}

		for y := range grid {
			for x := range grid[y] {
				newGrid[y][x] = grid[y][x]
				if grid[y][x] == "@" && checkPos(grid, x, y) {
					newGrid[y][x] = "."
					removed++
				}
			}
		}
		count += removed
		grid = newGrid
	}

	fmt.Println(count)
}
