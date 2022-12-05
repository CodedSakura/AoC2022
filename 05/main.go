package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

func A() {
	var setup []string
	isSetup := true

	var stacks [][]rune
	re, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)

	for line := range utils.ReadDayByLine(05) {
		if isSetup {
			if line == "" {
				isSetup = false

				for i := len(setup)/2 - 1; i >= 0; i-- {
					opp := len(setup) - 1 - i
					setup[i], setup[opp] = setup[opp], setup[i]
				}
				width := len(strings.Split(strings.Trim(setup[0], " "), "  "))
				_, setup = setup[0], setup[1:]

				stacks = make([][]rune, width)

				for _, stackLine := range setup {
					for i := 1; i < len(stackLine); i += 4 {
						if rune(stackLine[i]) == ' ' {
							continue
						}
						stacks[(i-1)/4] = append(stacks[(i-1)/4], rune(stackLine[i]))
					}
				}

				continue
			}

			setup = append(setup, line)
			continue
		}

		match := re.FindStringSubmatch(line)[1:]
		move, _ := strconv.Atoi(match[0])
		from, _ := strconv.Atoi(match[1])
		to, _ := strconv.Atoi(match[2])

		from--
		to--

		for i := 0; i < move; i++ {
			var item rune
			fromIndex := len(stacks[from]) - 1
			item, stacks[from] = stacks[from][fromIndex], stacks[from][:fromIndex]
			stacks[to] = append(stacks[to], item)
		}
	}

	for _, row := range stacks {
		fmt.Print(string(row[len(row)-1]))
	}
	fmt.Println()
}

func B() {
	var setup []string
	isSetup := true

	var stacks [][]rune
	re, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)

	for line := range utils.ReadDayByLine(05) {
		if isSetup {
			if line == "" {
				isSetup = false

				for i := len(setup)/2 - 1; i >= 0; i-- {
					opp := len(setup) - 1 - i
					setup[i], setup[opp] = setup[opp], setup[i]
				}
				width := len(strings.Split(strings.Trim(setup[0], " "), "  "))
				_, setup = setup[0], setup[1:]

				stacks = make([][]rune, width)

				for _, stackLine := range setup {
					for i := 1; i < len(stackLine); i += 4 {
						if rune(stackLine[i]) == ' ' {
							continue
						}
						stacks[(i-1)/4] = append(stacks[(i-1)/4], rune(stackLine[i]))
					}
				}

				continue
			}

			setup = append(setup, line)
			continue
		}

		match := re.FindStringSubmatch(line)[1:]
		move, _ := strconv.Atoi(match[0])
		from, _ := strconv.Atoi(match[1])
		to, _ := strconv.Atoi(match[2])

		from--
		to--

		var items []rune
		fromIndex := len(stacks[from]) - move
		items, stacks[from] = stacks[from][fromIndex:], stacks[from][:fromIndex]
		stacks[to] = append(stacks[to], items...)
	}

	for _, row := range stacks {
		fmt.Print(string(row[len(row)-1]))
	}
	fmt.Println()
}
