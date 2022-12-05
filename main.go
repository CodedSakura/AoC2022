package main

import (
	"AoC2022/utils"
	"fmt"
)

func main() {
	A()
	// Template file
}

func A() {
	result := 0

	for line := range utils.ReadDayByLine(00) {
		fmt.Println(line)
	}

	fmt.Println(result)
}
