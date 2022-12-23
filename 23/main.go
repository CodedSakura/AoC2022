package main

import (
	"AoC2022/utils"
	"fmt"
	"math"
)

const visA = false
const visB = false

func main() {
	if !visB {
		A()
	}
	if !visA {
		B()
	}
}

const (
	N = 1
	S = 2
	W = 3
	E = 4
)

type elf struct {
	proposedDir int
}

func getBounds(elves [][]*elf) [4]int {
	bounds := [4]int{
		math.MaxInt, math.MinInt, // n s
		math.MaxInt, math.MinInt, // w e
	}
	for y, row := range elves {
		for x, e := range row {
			if e == nil {
				continue
			}

			if y < bounds[0] {
				bounds[0] = y
			}
			if y > bounds[1] {
				bounds[1] = y
			}
			if x < bounds[2] {
				bounds[2] = x
			}
			if x > bounds[3] {
				bounds[3] = x
			}
		}
	}
	return bounds
}

func printElves(elves [][]*elf) {
	bounds := getBounds(elves)
	fmt.Println(bounds, bounds[1]-bounds[0], bounds[3]-bounds[2])
	for y := bounds[0]; y <= bounds[1]; y++ {
		for x := bounds[2]; x <= bounds[3]; x++ {
			if elves[y][x] == nil {
				fmt.Print(".")
			} else {
				switch elves[y][x].proposedDir {
				case N:
					fmt.Print("^")
					break
				case S:
					fmt.Print("v")
					break
				case W:
					fmt.Print("<")
					break
				case E:
					fmt.Print(">")
					break
				default:
					fmt.Print("#")
					break
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func count(elves [][]*elf) (elfCount, fieldCount int) {
	bounds := getBounds(elves)
	for y := bounds[0]; y <= bounds[1]; y++ {
		for x := bounds[2]; x <= bounds[3]; x++ {
			if elves[y][x] == nil {
				fieldCount++
				continue
			}
			elfCount++
		}
	}
	return
}

func adjustPos(x, y int, direction int) (int, int) {
	switch direction {
	case N:
		y--
		break
	case S:
		y++
		break
	case W:
		x--
		break
	case E:
		x++
		break
	}
	return x, y
}

func neighborsInDir(x, y int, elves [][]*elf, direction int) (out int) {
	var toCheck [3][2]int
	switch direction {
	case N:
		toCheck = [3][2]int{{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}}
		break
	case S:
		toCheck = [3][2]int{{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}}
		break
	case W:
		toCheck = [3][2]int{{x - 1, y - 1}, {x - 1, y}, {x - 1, y + 1}}
		break
	case E:
		toCheck = [3][2]int{{x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}}
		break
	}
	for _, pos := range toCheck {
		if elves[pos[1]][pos[0]] != nil {
			out++
		}
	}
	return
}

func getNeighbors(x, y int, elves [][]*elf) (out [3][3]*elf) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			out[dy+1][dx+1] = elves[y+dy][x+dx]
		}
	}
	return
}

func countMovingTo(x, y int, elves [][]*elf) (out int) {
	n := getNeighbors(x, y, elves)
	if n[0][1] != nil && n[0][1].proposedDir == S {
		out++
	}
	if n[2][1] != nil && n[2][1].proposedDir == N {
		out++
	}
	if n[1][0] != nil && n[1][0].proposedDir == E {
		out++
	}
	if n[1][2] != nil && n[1][2].proposedDir == W {
		out++
	}
	return
}
func clearMovingTo(x, y int, elves *[][]*elf) {
	n := getNeighbors(x, y, *elves)
	if n[0][1] != nil && n[0][1].proposedDir == S {
		(*elves)[y-1][x].proposedDir = 0
	}
	if n[2][1] != nil && n[2][1].proposedDir == N {
		(*elves)[y+1][x].proposedDir = 0
	}
	if n[1][0] != nil && n[1][0].proposedDir == E {
		(*elves)[y][x-1].proposedDir = 0
	}
	if n[1][2] != nil && n[1][2].proposedDir == W {
		(*elves)[y][x+1].proposedDir = 0
	}
}

func A() {
	var elves [][]*elf

	for line := range utils.ReadDayByLine(23) {
		currLine := len(elves)
		elves = append(elves, make([]*elf, len(line)*3))
		for i, c := range line {
			if c == '#' {
				elves[currLine][len(line)+i] = &elf{}
			}
		}
	}
	rowCount := len(elves)
	rowLen := len(elves[0])
	for i := 0; i < rowCount; i++ {
		elves = append([][]*elf{make([]*elf, rowLen)}, append(elves, make([]*elf, rowLen))...)
	}

	if visA {
		fmt.Printf("%d %d\n\n", len(elves), len(elves[0]))
	}

	for i := 0; i < 10; i++ {
		if visA {
			fmt.Println("step", i)
			printElves(elves)
		}

		directions := [4]int{(i)%4 + 1, (i+1)%4 + 1, (i+2)%4 + 1, (i+3)%4 + 1}
		hasProposedMoving := false

		for y, row := range elves {
			for x, e := range row {
				if e == nil {
					continue
				}

				prop := -1
				alone := true
				for _, direction := range directions {
					n := neighborsInDir(x, y, elves, direction)
					if n == 0 {
						if prop == -1 {
							prop = direction
						}
					} else {
						alone = false
					}
				}
				if !alone && prop != -1 {
					e.proposedDir = prop
					hasProposedMoving = true
				}
			}
		}
		if !hasProposedMoving {
			break
		}

		if visA {
			printElves(elves)
		}
		for y, row := range elves {
			for x, e := range row {
				if e == nil {
					continue
				}

				nx, ny := adjustPos(x, y, e.proposedDir)
				nc := countMovingTo(nx, ny, elves)
				e.proposedDir = 0
				if nc == 1 {
					elves[ny][nx] = e
					elves[y][x] = nil
				} else {
					clearMovingTo(nx, ny, &elves)
				}
			}
		}
	}
	if visA {
		printElves(elves)
	}

	_, f := count(elves)
	fmt.Println(f)
}

func B() {
	var elves [][]*elf

	for line := range utils.ReadDayByLine(23) {
		currLine := len(elves)
		elves = append(elves, make([]*elf, len(line)*3))
		for i, c := range line {
			if c == '#' {
				elves[currLine][len(line)+i] = &elf{}
			}
		}
	}
	rowCount := len(elves)
	rowLen := len(elves[0])
	for i := 0; i < rowCount; i++ {
		elves = append([][]*elf{make([]*elf, rowLen)}, append(elves, make([]*elf, rowLen))...)
	}

	if visB {
		fmt.Printf("%d %d\n\n", len(elves), len(elves[0]))
	}

	i := 0
	for {
		if visB {
			fmt.Println("step", i)
			printElves(elves)
		}

		directions := [4]int{(i)%4 + 1, (i+1)%4 + 1, (i+2)%4 + 1, (i+3)%4 + 1}
		hasProposedMoving := false

		for y, row := range elves {
			for x, e := range row {
				if e == nil {
					continue
				}

				prop := -1
				alone := true
				for _, direction := range directions {
					n := neighborsInDir(x, y, elves, direction)
					if n == 0 {
						if prop == -1 {
							prop = direction
						}
					} else {
						alone = false
					}
				}
				if !alone && prop != -1 {
					e.proposedDir = prop
					hasProposedMoving = true
				}
			}
		}
		if !hasProposedMoving {
			break
		}

		if visB {
			printElves(elves)
		}
		for y, row := range elves {
			for x, e := range row {
				if e == nil {
					continue
				}

				nx, ny := adjustPos(x, y, e.proposedDir)
				nc := countMovingTo(nx, ny, elves)
				e.proposedDir = 0
				if nc == 1 {
					elves[ny][nx] = e
					elves[y][x] = nil
				} else {
					clearMovingTo(nx, ny, &elves)
				}
			}
		}

		i++
	}
	if visB {
		printElves(elves)
	}

	fmt.Println(i + 1)
}
