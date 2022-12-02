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

const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
)

func resolve(opponent int, me int) int {
	switch opponent {
	case Rock:
		switch me {
		case Rock:
			return 0
		case Paper:
			return 1
		case Scissors:
			return -1
		}
	case Paper:
		switch me {
		case Rock:
			return -1
		case Paper:
			return 0
		case Scissors:
			return 1
		}
	case Scissors:
		switch me {
		case Rock:
			return 1
		case Paper:
			return -1
		case Scissors:
			return 0
		}
	}
	return 0
}

func A() {
	result := 0

	for line := range utils.ReadDayByLine(2) {
		split := strings.Split(line, " ")

		var opChoice int
		switch split[0] {
		case "A":
			opChoice = Rock
			break
		case "B":
			opChoice = Paper
			break
		case "C":
			opChoice = Scissors
			break
		}

		var myChoice int
		switch split[1] {
		case "X":
			myChoice = Rock
			break
		case "Y":
			myChoice = Paper
			break
		case "Z":
			myChoice = Scissors
			break
		}

		result += (3 + 3*resolve(opChoice, myChoice)) + myChoice
	}

	fmt.Println(result)
}

func B() {
	result := 0

	for line := range utils.ReadDayByLine(2) {
		split := strings.Split(line, " ")

		var opChoice int
		switch split[0] {
		case "A":
			opChoice = Rock
			break
		case "B":
			opChoice = Paper
			break
		case "C":
			opChoice = Scissors
			break
		}

		switch split[1] {
		case "X":
			result += ((opChoice + 1) % 3) + 1
			break
		case "Y":
			result += opChoice + 3
			break
		case "Z":
			result += (opChoice % 3) + 7
			break
		}
	}

	fmt.Println(result)
}
