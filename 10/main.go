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
	x := 1
	cycles := 0

	result := 0
	for line := range utils.ReadDayByLine(10) {
		cycles += 1
		if cycles%40 == 20 && cycles < 240 {
			result += cycles * x
		}

		if strings.HasPrefix(line, "addx") {
			cycles += 1

			if cycles%40 == 20 && cycles < 240 {
				result += cycles * x
			}

			val, _ := strconv.Atoi(line[5:])
			x += val
		} else if line != "noop" {
			panic(line)
		}
	}

	fmt.Println(result)
}

func drawSprite(x int, cycle int) string {
	xPos := cycle % 40

	newLine := ""
	if cycle%40 == 0 {
		newLine += "\n"
	}

	if x <= xPos && xPos <= x+2 {
		return "#" + newLine
	} else {
		return "." + newLine
	}
}
func B() {
	x := 1
	cycles := 0
	result := ""

	for line := range utils.ReadDayByLine(10) {
		cycles += 1
		result += drawSprite(x, cycles)

		if strings.HasPrefix(line, "addx") {
			cycles += 1
			result += drawSprite(x, cycles)

			val, _ := strconv.Atoi(line[5:])
			x += val
		} else if line != "noop" {
			panic(line)
		}
	}

	fmt.Println(result)
}
