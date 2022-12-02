package main

import (
	"AoC2022/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	A()
	B()
}

func A() {
	maxCarry := 0
	currCary := 0

	for line := range utils.ReadDayByLine(1) {
		if line == "" {
			if currCary > maxCarry {
				maxCarry = currCary
			}
			currCary = 0
			continue
		}

		num, _ := strconv.Atoi(line)
		currCary += num
	}

	if currCary > maxCarry {
		maxCarry = currCary
	}

	fmt.Println(maxCarry)
}

func B() {
	var sums []int
	currCary := 0

	for line := range utils.ReadDayByLine(1) {
		if line == "" {
			sums = append(sums, currCary)
			currCary = 0
			continue
		}

		num, _ := strconv.Atoi(line)
		currCary += num
	}
	sums = append(sums, currCary)

	sort.Ints(sums)

	res := 0
	for _, i := range sums[len(sums)-3:] {
		res += i
	}

	fmt.Println(res)
}
