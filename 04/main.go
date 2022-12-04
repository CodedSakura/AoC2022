package main

import (
	"AoC2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

func A() {
	result := 0

	for line := range utils.ReadDayByLine(04) {
		commaSplit := strings.Split(line, ",")
		rangeA, rangeB := commaSplit[0], commaSplit[1]
		rangeASplit := strings.Split(rangeA, "-")
		rangeBSplit := strings.Split(rangeB, "-")
		a1, _ := strconv.Atoi(rangeASplit[0])
		a2, _ := strconv.Atoi(rangeASplit[1])
		b1, _ := strconv.Atoi(rangeBSplit[0])
		b2, _ := strconv.Atoi(rangeBSplit[1])

		if (a1 >= b1 && a2 <= b2) || (a1 <= b1 && a2 >= b2) {
			result++
		}
	}

	fmt.Println(result)
}

func B() {
	result := 0

	for line := range utils.ReadDayByLine(04) {
		commaSplit := strings.Split(line, ",")
		rangeA, rangeB := commaSplit[0], commaSplit[1]
		rangeASplit := strings.Split(rangeA, "-")
		rangeBSplit := strings.Split(rangeB, "-")
		a1, _ := strconv.Atoi(rangeASplit[0])
		a2, _ := strconv.Atoi(rangeASplit[1])
		b1, _ := strconv.Atoi(rangeBSplit[0])
		b2, _ := strconv.Atoi(rangeBSplit[1])

		if (a1 <= b1 && b1 <= a2) || (a1 <= b2 && b2 <= a2) ||
			(b1 <= a1 && a1 <= b2) || (b1 <= a2 && a2 <= b2) {
			result++
		}
	}

	fmt.Println(result)
}
