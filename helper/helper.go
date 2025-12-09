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

func RealLineIntersect2d(l1p1, l1p2, l2p1, l2p2 [2]int) bool {
	horizontalL1 := l1p1[1] == l1p2[1]
	horizontalL2 := l2p1[1] == l2p2[1]
	if horizontalL1 == horizontalL2 {
		return false
	}
	if horizontalL1 {
		l2x := l2p1[0]
		if (l1p1[0] <= l2x && l1p2[0] <= l2x) || (l1p1[0] >= l2x && l1p2[0] >= l2x) {
			return false
		}
		return (l2p1[1]-l1p1[1])*(l2p2[1]-l1p1[1]) < 0
	}
	if horizontalL2 {
		l1x := l1p1[0]
		if (l2p1[0] <= l1x && l2p2[0] <= l1x) || (l2p1[0] >= l1x && l2p2[0] >= l1x) {
			return false
		}
		return (l1p1[1]-l2p1[1])*(l1p2[1]-l2p1[1]) < 0
	}
	panic("unreachable")
	return false
}

func RectLineIntersect2d(b1, b2, l1, l2 [2]int) bool {
	p1 := b1
	p2 := b2
	p3 := [2]int{b1[0], b2[1]}
	p4 := [2]int{b2[0], b1[1]}
	if RealLineIntersect2d(p1, p3, l1, l2) || RealLineIntersect2d(p2, p4, l1, l2) || RealLineIntersect2d(p1, p4, l1, l2) || RealLineIntersect2d(p2, p3, l1, l2) {
		return true
	}
	return false
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsInBounds2d(b1, b2, p [2]int) bool {
	if b2[0] > b1[0] && (p[0] <= b1[0] || p[0] >= b2[0]) {
		return false
	} else if b2[0] < b1[0] && (p[0] <= b2[0] || p[0] >= b1[0]) {
		return false
	}
	if b2[1] > b1[1] && (p[1] <= b1[1] || p[1] >= b2[1]) {
		return false
	} else if b2[1] < b1[1] && (p[1] <= b2[1] || p[1] >= b1[1]) {
		return false
	}
	return true
}

func IsInBounds2d2(b1, b2, p [2]int) bool {
	xIn := (b1[0]-p[0] < 0) != (b2[0]-p[0] < 0)
	yIn := (b1[1]-p[1] < 0) != (b2[1]-p[1] < 0)
	return xIn && yIn
}

func RectLineIntersect2d2(b1, b2, l1, l2 [2]int) bool {
	xIn1 := (b1[0]-l1[0] <= 0) != (b2[0]-l1[0] <= 0)
	xIn2 := (b1[0]-l2[0] <= 0) != (b2[0]-l2[0] <= 0)
	yIn1 := (b1[1]-l1[1] <= 0) != (b2[1]-l1[1] <= 0)
	yIn2 := (b1[1]-l2[1] <= 0) != (b2[1]-l2[1] <= 0)
	return (xIn1 != xIn2) != (yIn1 != yIn2)
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
