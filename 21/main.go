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

type yells struct {
	isNum    bool
	num      int
	opA, opB string
	operand  string
}

func A() {
	monkeys := make(map[string]yells)

	for line := range utils.ReadDayByLine(21) {
		split := strings.Split(line, ": ")
		if n, e := strconv.Atoi(split[1]); e == nil {
			monkeys[split[0]] = yells{
				isNum: true,
				num:   n,
			}
		} else {
			monkeys[split[0]] = yells{
				isNum:   false,
				opA:     split[1][:4],
				opB:     split[1][7:11],
				operand: split[1][5:6],
			}
		}
	}

	var getVal func(monkey string) int
	getVal = func(monkey string) int {
		if monkeys[monkey].isNum {
			return monkeys[monkey].num
		} else {
			mA := monkeys[monkey].opA
			mB := monkeys[monkey].opB

			vA := getVal(mA)
			vB := getVal(mB)

			switch monkeys[monkey].operand {
			case "+":
				return vA + vB
			case "-":
				return vA - vB
			case "*":
				return vA * vB
			case "/":
				return vA / vB
			}
		}
		return 0
	}

	fmt.Println(getVal("root"))
}

func B() {
	monkeys := make(map[string]yells)

	for line := range utils.ReadDayByLine(21) {
		split := strings.Split(line, ": ")
		if n, e := strconv.Atoi(split[1]); e == nil {
			monkeys[split[0]] = yells{
				isNum: true,
				num:   n,
			}
		} else {
			monkeys[split[0]] = yells{
				isNum:   false,
				opA:     split[1][:4],
				opB:     split[1][7:11],
				operand: split[1][5:6],
			}
		}
	}

	var getVal func(monkey string) int
	getVal = func(monkey string) int {
		if monkeys[monkey].isNum {
			return monkeys[monkey].num
		} else {
			mA := monkeys[monkey].opA
			mB := monkeys[monkey].opB

			vA := getVal(mA)
			vB := getVal(mB)

			switch monkeys[monkey].operand {
			case "+":
				return vA + vB
			case "-":
				return vA - vB
			case "*":
				return vA * vB
			case "/":
				return vA / vB
			case "=":
				return vA - vB
			}
		}
		return 0
	}

	root := monkeys["root"]
	root.operand = "="
	monkeys["root"] = root

	guessNum := func(num int) int {
		humn := yells{
			isNum: true,
			num:   num,
		}
		monkeys["humn"] = humn

		return getVal("root")
	}

	guess := 1
	magnitude := 0
	for {
		r := guessNum(guess)
		if r < 0 {
			break
		}
		magnitude++
		guess <<= 1
	}
	magnitude--

	for {
		res := guessNum(guess)
		//fmt.Println(guess, res)

		if res == 0 {
			// :upside-down:
			fmt.Println(guess - 1)
			return
		}

		if res > 0 {
			guess += 1 << magnitude
		} else {
			guess -= 1 << magnitude
		}
		magnitude--
	}
}
