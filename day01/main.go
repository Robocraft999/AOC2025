package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(1)
	//input = helper.TestInput(1, "testinput.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Fields(input)
	dial := 50
	res := 0
	for _, line := range lines {
		d, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		diff := int(d)
		if line[0] == 'L' {
			diff = 100 - diff
		}
		dial += diff
		dial %= 100

		if dial == 0 {
			res += 1
		}
	}
	fmt.Println(res)
}

func part2(input string) {
	lines := strings.Fields(input)
	dial := 50
	res := 0
	for _, line := range lines {
		d, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		diff := int(d)
		res += diff / 100
		diff %= 100
		if line[0] == 'L' {
			if dial-diff > 0 {
				diff *= -1
			} else {
				if dial == 0 || dial-diff == 0 {
					diff = 100 - diff
				} else {
					diff = 200 - diff
				}
			}
		}
		dial += diff
		//fmt.Println(line, dial, dial/100)
		res += dial / 100
		dial %= 100
	}
	fmt.Println(res)
}
