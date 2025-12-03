package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(3)
	input = helper.TestInput(3, "input.txt")
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		left := 0
		for i := 0; i < len(line); i++ {
			leftNum, _ := strconv.Atoi(line[left : left+1])
			leftL, _ := strconv.Atoi(line[i : i+1])
			if leftL > leftNum {
				left = i
			}
		}
		right := left + 1
		if right == len(line) {
			right = 0
		}
		for i := right; i < len(line); i++ {
			if i == left {
				continue
			}
			rightNum, _ := strconv.Atoi(line[right : right+1])
			rightL, _ := strconv.Atoi(line[i : i+1])
			if rightL > rightNum {
				right = i
			}
		}
		if left > right {
			temp := left
			left = right
			right = temp
		}
		numS := fmt.Sprintf("%s%s", line[left:left+1], line[right:right+1])
		num, _ := strconv.Atoi(numS)
		sum += num
	}
	fmt.Println(sum)
}

func getBiggestNumberLeftIndex(line string) int {
	biggestIndex := 0
	for i := 0; i < len(line); i++ {
		digit, _ := strconv.Atoi(string(line[i]))
		current, _ := strconv.Atoi(string(line[biggestIndex]))
		if digit > current {
			biggestIndex = i
		}
	}
	return biggestIndex
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		biggest := 0
		count := 12
		for i := 0; i < len(line); i++ {
			if i > len(line)-count {
				break
			} else if i == len(line)-count {
				num, _ := strconv.Atoi(line[i:])
				if num > biggest {
					biggest = num
				}
			}
			remLine := line[i:]
			resLine := ""
			for remaining := count; remaining > 0; remaining-- {
				nextBiggest := getBiggestNumberLeftIndex(remLine[:1+len(remLine)-remaining])
				resLine += string(remLine[nextBiggest])
				remLine = remLine[nextBiggest+1:]
			}
			num, _ := strconv.Atoi(resLine)
			if num > biggest {
				biggest = num
			}
		}
		sum += biggest
	}
	fmt.Println(sum)
}
