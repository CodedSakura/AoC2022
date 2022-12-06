package main

import (
	"AoC2022/utils"
	"fmt"
)

func main() {
	A()
	B()
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func A() {
	result := 0

	buffer := ""

	for line := range utils.ReadDayByLine(06) {
		for i, c := range line {
			buffer += string(c)
			buffer = buffer[max(len(buffer)-4, 0):]

			uniq := make(map[rune]struct{})
			for _, c2 := range buffer {
				uniq[c2] = struct{}{}
			}

			if len(uniq) == 4 {
				result = i + 1
				break
			}
		}
	}

	fmt.Println(result)
}

func B() {
	result := 0

	buffer := ""

	for line := range utils.ReadDayByLine(06) {
		for i, c := range line {
			buffer += string(c)
			buffer = buffer[max(len(buffer)-14, 0):]

			uniq := make(map[rune]struct{})
			for _, c2 := range buffer {
				uniq[c2] = struct{}{}
			}

			if len(uniq) == 14 {
				result = i + 1
				break
			}
		}
	}

	fmt.Println(result)
}
