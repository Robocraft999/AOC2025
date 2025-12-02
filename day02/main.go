package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(2)
	//input = helper.TestInput(2, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, rng := range ranges {
		parts := strings.Split(rng, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		halfS := int(math.Ceil(float64(len(parts[0])) / 2.0))
		halfE := int(math.Ceil(float64(len(parts[1])) / 2.0))
		for num := start; num <= end; num++ {
			sNum := strconv.Itoa(num)
			if len(sNum)%2 != 0 {
				continue
			}
			if strings.Count(sNum, sNum[:halfS]) == 2 {
				sum += num
			} else if strings.Count(sNum, sNum[:halfE]) == 2 {
				sum += num
			}
		}
	}
	fmt.Println(sum)
}

func part2(input string) {
	ranges := strings.Split(input, ",")
	sum := 0
	for _, rng := range ranges {
		parts := strings.Split(rng, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for num := start; num <= end; num++ {
			sNum := strconv.Itoa(num)
			for i := 1; i <= len(sNum)/2; i++ {
				if strings.Count(sNum, sNum[:i])*len(sNum[:i]) == len(sNum) {
					sum += num
					break
				}
			}
		}
	}
	fmt.Println(sum)
}
