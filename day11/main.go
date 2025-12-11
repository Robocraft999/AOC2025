package main

import (
	"fmt"
	"strings"

	"example.com/helper"
)

func main() {
	input := helper.Input(11)
	//input = helper.TestInput(11, "testinput.txt")
	part1(input)
	//input = helper.TestInput(11, "test2input.txt")
	part2(input)
}

func countPaths(current string, nodes map[string][]string, counts map[string]int, target string) int {
	if current == target {
		return 1
	}
	if _, ok := counts[current]; ok {
		return counts[current]
	}
	sum := 0
	for _, neighbor := range nodes[current] {
		sum += countPaths(neighbor, nodes, counts, target)
	}
	counts[current] = sum
	return sum
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	nodes := make(map[string][]string)
	counts := make(map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		start := strings.TrimSpace(parts[0])
		targets := strings.Split(strings.TrimSpace(parts[1]), " ")

		nodes[start] = targets
	}
	sum := countPaths("you", nodes, counts, "out")
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	nodes := make(map[string][]string)
	//paths := make(map[string]Node)
	//counts := make(map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		start := strings.TrimSpace(parts[0])
		targets := strings.Split(strings.TrimSpace(parts[1]), " ")

		nodes[start] = targets
	}
	{
		fftAmt := countPaths("svr", nodes, make(map[string]int), "fft")
		dacAmt := countPaths("fft", nodes, make(map[string]int), "dac")
		outAmt := countPaths("dac", nodes, make(map[string]int), "out")
		fmt.Println("fft first", fftAmt, dacAmt, outAmt, fftAmt*dacAmt*outAmt)
	}
	{
		dacAmt := countPaths("svr", nodes, make(map[string]int), "dac")
		fftAmt := countPaths("dac", nodes, make(map[string]int), "fft")
		outAmt := countPaths("fft", nodes, make(map[string]int), "out")
		fmt.Println("dac first", fftAmt, dacAmt, outAmt, fftAmt*dacAmt*outAmt)
	}
}
