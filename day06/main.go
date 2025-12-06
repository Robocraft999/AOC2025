package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(6)
	//input = helper.TestInput(6, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Fields(line)
	}
	helper.PrintGrid(grid)

	sum := 0
	for x := range grid[0] {
		columnRes := 0
		op := grid[len(grid)-1][x]
		for y := range len(grid) - 1 {
			if op == "*" {
				if columnRes == 0 {
					num, _ := strconv.Atoi(grid[y][x])
					columnRes = num
				} else {
					num, _ := strconv.Atoi(grid[y][x])
					columnRes *= num
				}
			} else if op == "+" {
				num, _ := strconv.Atoi(grid[y][x])
				columnRes += num
			}
		}
		sum += columnRes
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	grid := helper.MakeGridFromLines(lines)
	helper.PrintGrid(grid)

	sum := 0
	newColumn := true
	columnRes := 0
	op := ""
	for x := range grid[0] {
		if grid[0][x] == " " && x < len(grid[0])-1 && grid[len(grid)-1][x+1] != " " {
			newColumn = true
			continue
		}
		if newColumn {
			sum += columnRes
			columnRes = 0
			op = grid[len(grid)-1][x]
			newColumn = false
		}

		num := ""
		for y := range len(grid) - 1 {
			num += grid[y][x]
		}
		num = strings.TrimSpace(num)
		n, _ := strconv.Atoi(num)
		if op == "*" {
			if columnRes == 0 {
				columnRes = n
			} else {
				columnRes *= n
			}
		} else if op == "+" {
			columnRes += n
		}
	}
	sum += columnRes
	fmt.Println(sum)
}
