package main

import (
	"AoC2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

type operation struct {
	isMultiplication bool
	value            int
}

type test struct {
	isDivisibleBy int
	ifTrue        int
	ifFalse       int
}

type monkey struct {
	items []int
	operation
	test
	inspectedCount int
}

func processRound(monkeys []monkey, manageable bool, lcm int) {
	for i := range monkeys {
		monk := &monkeys[i]
		for len(monk.items) > 0 {
			item := monk.items[0]
			monk.items = monk.items[1:]

			if monk.operation.isMultiplication {
				if monk.operation.value == 0 {
					item *= item
				} else {
					item *= monk.operation.value
				}
			} else {
				item += monk.operation.value
			}

			if manageable {
				item /= 3
			} else {
				item = item % lcm
			}

			if item%monk.test.isDivisibleBy == 0 {
				passTo := &monkeys[monk.test.ifTrue]
				passTo.items = append(passTo.items, item)
			} else {
				passTo := &monkeys[monk.test.ifFalse]
				passTo.items = append(passTo.items, item)
			}

			monk.inspectedCount++
		}
	}
}

func A() {
	var monkeys []monkey

	for line := range utils.ReadDayByLine(11) {
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, monkey{})
		} else if strings.HasPrefix(line, "  ") {
			line = strings.TrimPrefix(line, "  ")
			currMonkey := &monkeys[len(monkeys)-1]
			if strings.HasPrefix(line, "Starting items: ") {
				line = strings.TrimPrefix(line, "Starting items: ")
				items := strings.Split(line, ", ")
				for _, item := range items {
					m, _ := strconv.Atoi(item)
					currMonkey.items = append(currMonkey.items, m)
				}
			} else if strings.HasPrefix(line, "Operation: ") {
				line = strings.TrimPrefix(line, "Operation: ")
				if !strings.HasPrefix(line, "new = old ") {
					panic(line)
				}

				line = strings.TrimPrefix(line, "new = old ")
				if line[0] == '*' {
					if line[2:] == "old" {
						currMonkey.operation = operation{isMultiplication: true, value: 0}
					} else {
						m, _ := strconv.Atoi(line[2:])
						currMonkey.operation = operation{isMultiplication: true, value: m}
					}
				} else if line[0] == '+' {
					m, _ := strconv.Atoi(line[2:])
					currMonkey.operation = operation{isMultiplication: false, value: m}
				} else {
					panic(line)
				}
			} else if strings.HasPrefix(line, "Test: ") {
				line = strings.TrimPrefix(line, "Test: ")
				if !strings.HasPrefix(line, "divisible by") {
					panic(line)
				}
				line = strings.TrimPrefix(line, "divisible by ")
				m, _ := strconv.Atoi(line)
				currMonkey.test.isDivisibleBy = m
			} else if strings.HasPrefix(line, "  ") {
				line = strings.TrimPrefix(line, "  ")
				if strings.HasPrefix(line, "If true: ") {
					line = strings.TrimPrefix(line, "If true: ")
					if !strings.HasPrefix(line, "throw to monkey ") {
						panic(line)
					}
					m, _ := strconv.Atoi(line[16:])
					currMonkey.test.ifTrue = m
				} else if strings.HasPrefix(line, "If false: ") {
					line = strings.TrimPrefix(line, "If false: ")
					if !strings.HasPrefix(line, "throw to monkey ") {
						panic(line)
					}
					m, _ := strconv.Atoi(line[16:])
					currMonkey.test.ifFalse = m
				} else {
					panic(line)
				}
			} else {
				panic(line)
			}
		}
	}

	for i := 0; i < 20; i++ {
		processRound(monkeys, true, 0)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	fmt.Println(monkeys[0].inspectedCount * monkeys[1].inspectedCount)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func B() {
	var monkeys []monkey

	for line := range utils.ReadDayByLine(11) {
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, monkey{})
		} else if strings.HasPrefix(line, "  ") {
			line = strings.TrimPrefix(line, "  ")
			currMonkey := &monkeys[len(monkeys)-1]
			if strings.HasPrefix(line, "Starting items: ") {
				line = strings.TrimPrefix(line, "Starting items: ")
				items := strings.Split(line, ", ")
				for _, item := range items {
					m, _ := strconv.Atoi(item)
					currMonkey.items = append(currMonkey.items, m)
				}
			} else if strings.HasPrefix(line, "Operation: ") {
				line = strings.TrimPrefix(line, "Operation: ")
				if !strings.HasPrefix(line, "new = old ") {
					panic(line)
				}

				line = strings.TrimPrefix(line, "new = old ")
				if line[0] == '*' {
					if line[2:] == "old" {
						currMonkey.operation = operation{isMultiplication: true, value: 0}
					} else {
						m, _ := strconv.Atoi(line[2:])
						currMonkey.operation = operation{isMultiplication: true, value: m}
					}
				} else if line[0] == '+' {
					m, _ := strconv.Atoi(line[2:])
					currMonkey.operation = operation{isMultiplication: false, value: m}
				} else {
					panic(line)
				}
			} else if strings.HasPrefix(line, "Test: ") {
				line = strings.TrimPrefix(line, "Test: ")
				if !strings.HasPrefix(line, "divisible by") {
					panic(line)
				}
				line = strings.TrimPrefix(line, "divisible by ")
				m, _ := strconv.Atoi(line)
				currMonkey.test.isDivisibleBy = m
			} else if strings.HasPrefix(line, "  ") {
				line = strings.TrimPrefix(line, "  ")
				if strings.HasPrefix(line, "If true: ") {
					line = strings.TrimPrefix(line, "If true: ")
					if !strings.HasPrefix(line, "throw to monkey ") {
						panic(line)
					}
					m, _ := strconv.Atoi(line[16:])
					currMonkey.test.ifTrue = m
				} else if strings.HasPrefix(line, "If false: ") {
					line = strings.TrimPrefix(line, "If false: ")
					if !strings.HasPrefix(line, "throw to monkey ") {
						panic(line)
					}
					m, _ := strconv.Atoi(line[16:])
					currMonkey.test.ifFalse = m
				} else {
					panic(line)
				}
			} else {
				panic(line)
			}
		}
	}

	divisorLCM := monkeys[0].test.isDivisibleBy
	for _, m := range monkeys {
		divisorLCM = lcm(divisorLCM, m.test.isDivisibleBy)
	}

	for i := 0; i < 10_000; i++ {
		processRound(monkeys, false, divisorLCM)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedCount > monkeys[j].inspectedCount
	})

	fmt.Println(monkeys[0].inspectedCount * monkeys[1].inspectedCount)
}
