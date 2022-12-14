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
	x, y int
}

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func gridPointsToGrid(points map[pos]struct{}, p2 bool) [][]int {
	maxX, maxY := -1, -1
	for p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	if p2 {
		maxY += 2
		maxX += maxY
	}

	grid := make([][]int, maxY+1)
	for y := 0; y < maxY+1; y++ {
		grid[y] = make([]int, maxX+1)
	}

	for p := range points {
		grid[p.y][p.x] = 1
	}

	if p2 {
		for i := 0; i < maxX+1; i++ {
			grid[maxY][i] = 1
		}
	}

	return grid
}

func dropSand(grid *[][]int) bool {
	x, y := 500, 0
	for {
		if y+1 >= len(*grid) {
			return false
		}

		if (*grid)[y+1][x] == 0 {
			y++
			continue
		}
		if (*grid)[y+1][x-1] == 0 {
			y++
			x--
			continue
		}
		if (*grid)[y+1][x+1] == 0 {
			y++
			x++
			continue
		}

		if x == 500 && y == 0 {
			return false
		}
		break
	}
	(*grid)[y][x] = 2
	return true
}

func A() {
	gridPoints := make(map[pos]struct{})

	for line := range utils.ReadDayByLine(14) {
		var ppos pos
		for _, strPos := range strings.Split(line, " -> ") {
			posStr := strings.Split(strPos, ",")
			x, _ := strconv.Atoi(posStr[0])
			y, _ := strconv.Atoi(posStr[1])
			if (ppos != pos{}) {
				if x != ppos.x {
					a, b := minMax(x, ppos.x)
					for i := a; i <= b; i++ {
						gridPoints[pos{x: i, y: y}] = struct{}{}
					}
				} else {
					a, b := minMax(y, ppos.y)
					for i := a; i <= b; i++ {
						gridPoints[pos{x: x, y: i}] = struct{}{}
					}
				}
			}
			ppos = pos{x, y}
		}
	}

	grid := gridPointsToGrid(gridPoints, false)
	count := 0
	for dropSand(&grid) {
		count++
	}

	fmt.Println(count)
}

func B() {
	gridPoints := make(map[pos]struct{})

	for line := range utils.ReadDayByLine(14) {
		var ppos pos
		for _, strPos := range strings.Split(line, " -> ") {
			posStr := strings.Split(strPos, ",")
			x, _ := strconv.Atoi(posStr[0])
			y, _ := strconv.Atoi(posStr[1])
			if (ppos != pos{}) {
				if x != ppos.x {
					a, b := minMax(x, ppos.x)
					for i := a; i <= b; i++ {
						gridPoints[pos{x: i, y: y}] = struct{}{}
					}
				} else {
					a, b := minMax(y, ppos.y)
					for i := a; i <= b; i++ {
						gridPoints[pos{x: x, y: i}] = struct{}{}
					}
				}
			}
			ppos = pos{x, y}
		}
	}

	grid := gridPointsToGrid(gridPoints, true)
	count := 0
	for dropSand(&grid) {
		count++
	}

	fmt.Println(count + 1)
}
