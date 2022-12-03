package main

import (
	"AoC2022/utils"
	"fmt"
	"strings"
)

func main() {
	A()
	B()
}

func A() {
	prioritySum := 0

	for line := range utils.ReadDayByLine(03) {
		middle := len(line) / 2

		otherCompartment := line[middle:]
		overlap := make(map[rune]bool)

		for _, ch := range line[0:middle] {
			if strings.IndexRune(otherCompartment, ch) >= 0 {
				overlap[ch] = true
			}
		}

		for v := range overlap {
			if 'a' <= v && v <= 'z' {
				prioritySum += int(v) - 'a' + 1
			} else {
				prioritySum += int(v) - 'A' + 26 + 1
			}
		}
	}

	fmt.Println(prioritySum)
}

func B() {
	prioritySum := 0

	var prevA string
	var prevB string
	counter := 0

	for line := range utils.ReadDayByLine(03) {
		if counter == 0 {
			prevA = line
			counter++
			continue
		} else if counter == 1 {
			prevB = line
			counter++
			continue
		}
		overlap := make(map[rune]bool)

		for _, ch := range prevA {
			if strings.IndexRune(prevB, ch) >= 0 {
				overlap[ch] = true
			}
		}

		for v := range overlap {
			if strings.IndexRune(line, v) < 0 {
				overlap[v] = false
			}
		}

		for v, e := range overlap {
			if !e {
				continue
			}

			if 'a' <= v && v <= 'z' {
				prioritySum += int(v) - 'a' + 1
			} else {
				prioritySum += int(v) - 'A' + 26 + 1
			}
		}

		counter = 0
	}

	fmt.Println(prioritySum)
}
