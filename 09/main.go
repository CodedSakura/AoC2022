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

type pos struct {
	x int
	y int
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func printMap(tailVisited map[pos]struct{}) {
	min := pos{}
	max := pos{}
	for p := range tailVisited {
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}

		if p.x < min.x {
			min.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
	}

	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			if _, hasKey := tailVisited[pos{x: x, y: y}]; hasKey {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func printRope(head pos, tails [9]pos) {
	min := pos{}
	max := pos{}
	for _, p := range append([]pos{head}, tails[0:9]...) {
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}

		if p.x < min.x {
			min.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
	}

	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			drawPos := pos{x: x, y: y}
			if drawPos == head {
				fmt.Print("H")
			} else {
				tailpiece := false
				for i, p := range tails {
					if drawPos == p {
						tailpiece = true
						fmt.Print(i + 1)
						break
					}
				}
				if !tailpiece {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveA(head *pos, tail *pos, vec pos, tailVisited map[pos]struct{}) {
	if abs(vec.x)+abs(vec.y) > 1 {
		for i := 0; i < abs(vec.x); i++ {
			moveA(head, tail, pos{x: sign(vec.x)}, tailVisited)
		}
		for i := 0; i < abs(vec.y); i++ {
			moveA(head, tail, pos{y: sign(vec.y)}, tailVisited)
		}
		return
	}
	head.x += vec.x
	head.y += vec.y
	if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
		// move tail
		if abs(head.x-tail.x) > 1 {
			if head.y != tail.y {
				tail.y = head.y
			}
			tail.x += head.x - tail.x - sign(head.x-tail.x)
		}
		if abs(head.y-tail.y) > 1 {
			if head.x != tail.x {
				tail.x = head.x
			}
			tail.y += head.y - tail.y - sign(head.y-tail.y)
		}
	}
	tailVisited[*tail] = struct{}{}
}

func A() {
	tailVisited := make(map[pos]struct{})
	tail := pos{x: 0, y: 0}
	head := pos{x: 0, y: 0}

	tailVisited[tail] = struct{}{}

	for line := range utils.ReadDayByLine(9) {
		res := strings.Split(line, " ")
		amount, _ := strconv.Atoi(res[1])
		vec := pos{}
		switch res[0] {
		case "R":
			vec = pos{x: amount}
			break
		case "L":
			vec = pos{x: -amount}
			break
		case "D":
			vec = pos{y: -amount}
			break
		case "U":
			vec = pos{y: amount}
		}

		moveA(&head, &tail, vec, tailVisited)
	}

	fmt.Println(len(tailVisited))
}

func adjust(head *pos, tail *pos) {
	if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
		if abs(head.x-tail.x) > 1 && abs(head.y-tail.y) > 1 {
			tail.x += head.x - tail.x - sign(head.x-tail.x)
			tail.y += head.y - tail.y - sign(head.y-tail.y)
			return
		}
		if abs(head.x-tail.x) > 1 {
			if head.y != tail.y {
				tail.y = head.y
			}
			tail.x += head.x - tail.x - sign(head.x-tail.x)
		}
		if abs(head.y-tail.y) > 1 {
			if head.x != tail.x {
				tail.x = head.x
			}
			tail.y += head.y - tail.y - sign(head.y-tail.y)
		}
	}
}

func moveB(head *pos, tails *[9]pos, vec pos, tailVisited map[pos]struct{}) {
	if abs(vec.x)+abs(vec.y) > 1 {
		for i := 0; i < abs(vec.x); i++ {
			moveB(head, tails, pos{x: sign(vec.x)}, tailVisited)
		}
		for i := 0; i < abs(vec.y); i++ {
			moveB(head, tails, pos{y: sign(vec.y)}, tailVisited)
		}
		return
	}
	head.x += vec.x
	head.y += vec.y
	adjust(head, &tails[0])
	//printRope(*head, *tails)
	for i := 0; i < 8; i++ {
		adjust(&tails[i], &tails[i+1])
		//printRope(*head, *tails)
	}
	//fmt.Println()
	tailVisited[tails[8]] = struct{}{}
}

func B() {
	tailVisited := make(map[pos]struct{})
	head := pos{x: 0, y: 0}
	tails := [9]pos{}

	tailVisited[pos{}] = struct{}{}

	for line := range utils.ReadDayByLine(9) {
		res := strings.Split(line, " ")
		amount, _ := strconv.Atoi(res[1])
		vec := pos{}
		switch res[0] {
		case "R":
			vec = pos{x: amount}
			break
		case "L":
			vec = pos{x: -amount}
			break
		case "D":
			vec = pos{y: -amount}
			break
		case "U":
			vec = pos{y: amount}
		}

		moveB(&head, &tails, vec, tailVisited)
		//printRope(head, tails)
	}

	//printMap(tailVisited)

	fmt.Println(len(tailVisited))
}
