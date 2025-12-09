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
	part2(input)
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
	for _, first := range coords {
		for _, second := range coords {
			if _, ok := sizeMap[[2][2]int{second, first}]; !ok {
				w := helper.AbsInt(second[0]-first[0]) + 1
				h := helper.AbsInt(second[1]-first[1]) + 1
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

	fmt.Println("Part 1:", sizeMap[keys[0]])
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
	for _, first := range coords {
		for _, second := range coords {
			if _, ok := sizeMap[[2][2]int{second, first}]; !ok {
				flag := false
				for k, third := range coords {
					if helper.IsInBounds2d(first, second, third) {
						flag = true
						break
					}

					fourth := coords[(k+1+len(coords))%len(coords)]
					if helper.RectLineIntersect2d(first, second, third, fourth) {
						flag = true
						break
					}
				}

				w := helper.AbsInt(second[0]-first[0]) + 1
				h := helper.AbsInt(second[1]-first[1]) + 1
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

	fmt.Println("Part 2:", keys[0], sizeMap[keys[0]])
}
