package helper

import (
	"fmt"
	"math"
	"strings"
)

func Test() {
	fmt.Println("Helper Test")
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func EuclideanDistance(pos1 [3]int, pos2 [3]int) float64 {
	a := math.Pow(float64(pos1[0]-pos2[0]), 2)
	b := math.Pow(float64(pos1[1]-pos2[1]), 2)
	c := math.Pow(float64(pos1[2]-pos2[2]), 2)
	return math.Sqrt(a + b + c)
}

func MakeGridFromLines(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

func CountNeighbours8(grid [][]string, x, y int, search string) int {
	count := 0
	height := len(grid)
	width := len(grid[0])
	if x > 0 && y > 0 && grid[y-1][x-1] == search {
		count++
	}
	if x > 0 && grid[y][x-1] == search {
		count++
	}
	if x > 0 && y < height-1 && grid[y+1][x-1] == search {
		count++
	}

	if y > 0 && grid[y-1][x] == search {
		count++
	}
	if y < height-1 && grid[y+1][x] == search {
		count++
	}

	if x < width-1 && y > 0 && grid[y-1][x+1] == search {
		count++
	}
	if x < width-1 && grid[y][x+1] == search {
		count++
	}
	if x < width-1 && y < height-1 && grid[y+1][x+1] == search {
		count++
	}
	return count
}

func CountNeighbours4(grid [][]string, x, y int, search string) int {
	count := 0
	height := len(grid)
	width := len(grid[0])

	if x > 0 && grid[y][x-1] == search {
		count++
	}
	if y > 0 && grid[y-1][x] == search {
		count++
	}
	if y < height-1 && grid[y+1][x] == search {
		count++
	}
	if x < width-1 && grid[y][x+1] == search {
		count++
	}
	return count
}
