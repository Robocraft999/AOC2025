package main

import (
	"fmt"
	sort2 "sort"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(9)
	//input = helper.TestInput(9, "testinput.txt")
	part1(input)
	part21(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	coords := make([][2]int, len(lines))
	for i, line := range lines {
		coordsS := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(coordsS[0])
		y, _ := strconv.Atoi(coordsS[1])
		coords[i] = [2]int{x, y}
	}

	sizeMap := make(map[[2][2]int]int)
	for _, coord := range coords {
		for _, coord2 := range coords {
			if _, ok := sizeMap[[2][2]int{coord2, coord}]; !ok {
				w := coord2[0] - coord[0]
				if w < 0 {
					w = -w
				}
				w += 1
				h := coord2[1] - coord[1]
				if h < 0 {
					h = -h
				}
				h += 1
				sizeMap[[2][2]int{coord, coord2}] = w * h
			}
		}
	}

	keys := make([][2][2]int, 0, len(sizeMap))

	for key := range sizeMap {
		keys = append(keys, key)
	}
	sort2.SliceStable(keys, func(i, j int) bool {
		return sizeMap[keys[i]] > sizeMap[keys[j]]
	})

	fmt.Println("Part 1:", sizeMap[keys[0]])
}

func IsOnEdge(b1, b2, p [2]int) (bool, bool, bool) {
	if p == b1 || p == b2 {
		return false, true, true
	}
	sameX := (p[0] == b1[0] || p[0] == b2[0]) && helper.AbsInt(p[1]-b2[1]) < helper.AbsInt(b1[1]-b2[1])
	sameY := (p[1] == b1[1] || p[1] == b2[1]) && helper.AbsInt(p[0]-b2[0]) < helper.AbsInt(b1[0]-b2[0])
	return sameX || sameY, sameX, sameY
}

func part21(input string) {
	lines := strings.Split(input, "\n")
	coords := make([][2]int, len(lines))
	for i, line := range lines {
		coordsS := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(coordsS[0])
		y, _ := strconv.Atoi(coordsS[1])
		coords[i] = [2]int{x, y}
	}

	sizeMap := make(map[[2][2]int]int)
	for _, first := range coords {
		for _, second := range coords {
			if _, ok := sizeMap[[2][2]int{second, first}]; !ok {
				flag := false
				for k, third := range coords {
					//fmt.Println(first, second, third, helper.IsInBounds2d(first, second, third))
					if helper.IsInBounds2d(first, second, third) {
						flag = true
						//fmt.Println("In Range", first, second, third)
						break
					}

					fourth := coords[(k+1+len(coords))%len(coords)]
					e1, sx1, sy1 := IsOnEdge(first, second, third)
					e2, sx2, sy2 := IsOnEdge(first, second, fourth)
					if e1 && e2 && ((sx1 && sx2 && third[0] != fourth[0]) || (sy1 && sy2 && third[1] != fourth[1])) {
						flag = true
						//fmt.Println("On Edge", first, second, third, fourth)
						break
					}
				}

				w := second[0] - first[0]
				if w < 0 {
					w = -w
				}
				w += 1
				h := second[1] - first[1]
				if h < 0 {
					h = -h
				}
				h += 1
				//fmt.Println(first, second, w*h)
				if flag {
					continue
				}

				sizeMap[[2][2]int{first, second}] = w * h
			}
		}
	}

	keys := make([][2][2]int, 0, len(sizeMap))

	for key := range sizeMap {
		keys = append(keys, key)
	}
	sort2.SliceStable(keys, func(i, j int) bool {
		return sizeMap[keys[i]] > sizeMap[keys[j]]
	})

	for _, key := range keys {
		fmt.Println(key, sizeMap[key])
	}

	fmt.Println("Part 21:", keys[0], sizeMap[keys[0]])
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	coords := make([][2]int, len(lines))
	for i, line := range lines {
		coordsS := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(coordsS[0])
		y, _ := strconv.Atoi(coordsS[1])
		coords[i] = [2]int{x, y}
	}

	sizeMap := make(map[[2][2]int]int)
	for i, first := range coords {
		before := coords[(i-1+len(coords))%len(coords)]
		middle := coords[(i+1+len(coords))%len(coords)]
		second := coords[(i+2+len(coords))%len(coords)]
		after := coords[(i+3+len(coords))%len(coords)]

		startIndex := i + len(coords) + 3
		endIndex := i + 2*len(coords)

		//outside
		if helper.IsInBounds2d(before, after, middle) {
			fmt.Println("outside", before, after, middle, first, second)
			continue
		}

		flag := false
		for j := startIndex; j < endIndex; j++ {
			if helper.IsInBounds2d(first, second, coords[j%len(coords)]) {
				flag = true
				fmt.Println("In Range", first, second, coords[j%len(coords)])
				break
			}
		}
		if flag {
			continue
		}

		w, h := helper.AbsInt(second[0]-first[0])+1, helper.AbsInt(second[1]-first[1])+1
		sizeMap[[2][2]int{first, second}] = w * h
	}

	keys := make([][2][2]int, 0, len(sizeMap))

	for key := range sizeMap {
		keys = append(keys, key)
	}
	sort2.SliceStable(keys, func(i, j int) bool {
		return sizeMap[keys[i]] > sizeMap[keys[j]]
	})

	/*for _, key := range keys {
		fmt.Println(key, sizeMap[key])
	}*/

	fmt.Println("Part 2:", sizeMap[keys[0]])
}
