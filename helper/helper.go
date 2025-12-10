package helper

import (
	"fmt"
	"math"
	"strings"

	"example.com/helper/set"
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

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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

func PathLength(start, end [2]int, validator func([2]int) bool) int {
	visited := set.NewEmptySet[[2]int]()
	toProcess := [][3]int{{start[0], start[1], 0}}
	for len(toProcess) > 0 {
		current := toProcess[0]
		toProcess = toProcess[1:]
		c, f := [2]int{current[0], current[1]}, current[2]
		if c == end {
			return f
		}
		dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range dirs {
			npx, npy := c[0]+dir[0], c[1]+dir[1]
			newPos := [2]int{npx, npy}
			if validator(newPos) {
				if !visited.Contains(newPos) {
					toProcess = append(toProcess, [3]int{newPos[0], newPos[1], f + 1})
				}
			}
		}
		visited.Add(c)
	}
	return -1
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
