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
	input = helper.TestInput(11, "test2input.txt")
	part2(input)
}

func countPaths(current string, nodes map[string][]string, counts map[string]int) int {
	if current == "out" {
		return 1
	}
	if _, ok := counts[current]; ok {
		return counts[current]
	}
	sum := 0
	for _, neighbor := range nodes[current] {
		sum += countPaths(neighbor, nodes, counts)
	}
	counts[current] = sum
	return sum
}

type Path []string
type Node struct {
	amtPaths int
	dac      bool
	fft      bool
}

/*
svr,aaa,fft,ccc,ddd,eee,dac,fff,ggg,out
svr,aaa,fft,ccc,ddd,eee,dac,fff,hhh,out
svr,aaa,fft,ccc,eee,dac,fff,ggg,out
svr,aaa,fft,ccc,eee,dac,fff,hhh,out
svr,bbb,tty,ccc,ddd,eee,dac,fff,ggg,out
svr,bbb,tty,ccc,ddd,eee,dac,fff,hhh,out
svr,bbb,tty,ccc,eee,dac,fff,ggg,out
svr,bbb,tty,ccc,eee,dac,fff,hhh,out
*/

func traversePaths(current string, currentPath Path, nodes map[string][]string, paths map[string]Node, sum *int) {
	//dac := slices.Contains(currentPath, "dac")
	//fft := slices.Contains(currentPath, "fft")
	cn := Node{amtPaths: 0, dac: false, fft: false}
	if current == "dac" {
		cn.dac = true
	}
	if current == "fft" {
		cn.fft = true
	}

	if n, ok := paths[current]; ok {
		cn.dac = cn.dac || n.dac
		cn.fft = cn.fft || n.fft
		return
	}

	summ := 0
	for _, neighbor := range nodes[current] {
		if neighbor == "out" {
			summ += 1
			continue
		}
		traversePaths(neighbor, append(currentPath, current), nodes, paths, sum)
		summ += paths[neighbor].amtPaths
	}
	cn.amtPaths = summ
	paths[current] = cn
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
	sum := countPaths("you", nodes, counts)
	fmt.Println(sum)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	nodes := make(map[string][]string)
	paths := make(map[string]Node)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		start := strings.TrimSpace(parts[0])
		targets := strings.Split(strings.TrimSpace(parts[1]), " ")

		nodes[start] = targets
	}
	sum := 0
	traversePaths("svr", Path{}, nodes, paths, &sum)
	fmt.Println(sum)
}
