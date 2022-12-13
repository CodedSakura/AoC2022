package main

import (
	"AoC2022/utils"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

type item struct {
	isList bool
	value  int
	values []item
}

func valToList(val item) item {
	return item{isList: true, values: []item{val}}
}

func check(left, right item) int {
	if !left.isList && !right.isList {
		if left.value < right.value {
			return 1
		}
		if left.value > right.value {
			return -1
		}
	} else {
		if left.isList != right.isList {
			if left.isList {
				right = valToList(right)
			} else {
				left = valToList(left)
			}
		}

		lVal, rVal := left.values, right.values
		for i := 0; i < len(lVal); i++ {
			if i >= len(rVal) {
				return -1
			}

			res := check(lVal[i], rVal[i])
			if res != 0 {
				return res
			}
		}
		if len(lVal) < len(rVal) {
			return 1
		}
	}

	return 0
}

func findClosingBrace(line string, index int) int {
	depth := 0
	for i, c := range line[index:] {
		if c == '[' {
			depth++
		} else if c == ']' {
			depth--
		}

		if depth == 0 {
			return i + index
		}
	}
	return -1
}

func parseLine(line string) item {
	out := item{isList: true}

	line = line[1 : len(line)-1]

	for len(line) > 0 {
		if line[0] == '[' {
			brace := findClosingBrace(line, 0)
			newItem := parseLine(line[:brace+1])
			out.values = append(out.values, newItem)
			line = line[brace+1:]
		} else if line[0] == ',' {
			line = line[1:]
		} else {
			values := strings.SplitN(line, ",", 2)
			newItem := item{isList: false}
			n, _ := strconv.Atoi(values[0])
			newItem.value = n
			out.values = append(out.values, newItem)
			if len(values) == 1 {
				line = ""
			} else {
				line = values[1]
			}
		}
	}

	return out
}

func A() {
	result := 0

	index := 1
	var firstLine, secondLine item
	readingFirstLine := true

	for line := range utils.ReadDayByLine(13) {
		if line == "" {
			//fmt.Printf("\n%v\n%v\n", firstLine, secondLine)
			if check(firstLine, secondLine) > 0 {
				//fmt.Println(index)
				result += index
			}
			index++
			readingFirstLine = true
			continue
		}
		if readingFirstLine {
			firstLine = parseLine(line)
			readingFirstLine = false
		} else {
			secondLine = parseLine(line)
		}
	}
	//fmt.Printf("\n%v\n%v\n", firstLine, secondLine)
	if check(firstLine, secondLine) > 0 {
		//fmt.Println(index)
		result += index
	}

	fmt.Println(result)
}

func B() {
	sepA := parseLine("[[2]]")
	sepB := parseLine("[[6]]")
	packets := []item{sepA, sepB}

	for line := range utils.ReadDayByLine(13) {
		if line == "" {
			continue
		}
		packets = append(packets, parseLine(line))
	}

	sort.Slice(packets, func(i, j int) bool {
		return check(packets[i], packets[j]) > 0
	})

	posA, posB := -1, -1
	for i, p := range packets {
		if reflect.DeepEqual(p, sepA) {
			posA = i + 1
		}
		if reflect.DeepEqual(p, sepB) {
			posB = i + 1
		}
	}
	fmt.Println(posA * posB)

}
