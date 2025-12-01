package main

import (
	"fmt"
	"strconv"
	"strings"
)
import "example.com/helper"

func main() {
	input := helper.Input(1)
	//input := helper.TestInput(1, "testinput.txt")
	lines := strings.Fields(input)
	dial := 50
	res := 0
	for _, line := range lines {
		diff, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		if line[0] == 'L' {
			diff *= -1
		}
		dial += int(diff)
		for dial < 0 {
			if dial-int(diff) != 0 {
				//fmt.Println("L", line, dial)
				res += 1
			}
			dial += 100
		}
		for dial >= 100 {
			if dial != 100 {
				//fmt.Println("R", line, dial)
				res += 1
			}
			dial -= 100
		}
		if dial == 0 {
			//fmt.Println("E", line, dial)
			res += 1
		}
	}
	fmt.Println(res)
}
