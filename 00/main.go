package main

import (
	. "AoC2022/utils"
	"fmt"
	"math"
)

func main() {
	B()
}

func A() {
	prevNum := math.MaxInt
	res := 0

	for num := range ReadDayByIntPerLine(0) {
		if prevNum < num {
			res++
		}
		prevNum = num
	}

	fmt.Println(res)
}

func B() {
	prevNum := math.MaxInt
	res := 0

	values := ChanToArr(ReadDayByIntPerLine(0))

	for i := 2; i < len(values); i++ {
		num := values[i] + values[i-1] + values[i-2]
		if num > prevNum {
			res++
		}
		prevNum = num
	}

	fmt.Println(res)
}
