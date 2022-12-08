package main

import (
	"AoC2022/utils"
	"fmt"
)

func main() {
	A()
	B()
}

func A() {
	var grid [][]int

	for line := range utils.ReadDayByLine(8) {
		grid = append(grid, make([]int, len(line)))
		for i, char := range line {
			grid[len(grid)-1][i] = int(char - '0')
		}
	}

	visible := make([][]bool, len(grid))
	for i := range grid {
		visible[i] = make([]bool, len(grid[i]))
	}

	for y, row := range grid {
		maxL := -1
		for x, val := range row {
			if val > maxL {
				visible[y][x] = true
				maxL = val
			}
		}

		maxR := -1
		for x := len(row) - 1; x >= 0; x-- {
			val := row[x]
			if val > maxR {
				visible[y][x] = true
				maxR = val
			}
		}
	}

	for x := 0; x < len(grid[0]); x++ {
		maxT := -1
		for y := 0; y < len(grid); y++ {
			val := grid[y][x]
			if val > maxT {
				visible[y][x] = true
				maxT = val
			}
		}

		maxB := -1
		for y := len(grid) - 1; y >= 0; y-- {
			val := grid[y][x]
			if val > maxB {
				visible[y][x] = true
				maxB = val
			}
		}
	}

	result := 0
	for _, row := range visible {
		for _, val := range row {
			if val {
				result++
			}
		}
	}

	fmt.Println(result)
}

func getTreeScore(x int, y int, grid [][]int) int {
	if x == 0 || y == 0 || x == len(grid[0])-1 || y == len(grid)-1 {
		return 0
	}

	up := 0
	down := 0
	left := 0
	right := 0

	for cy := y - 1; cy >= 0; cy-- {
		up++
		if grid[cy][x] >= grid[y][x] {
			break
		}
	}

	for cy := y + 1; cy < len(grid[0]); cy++ {
		down++
		if grid[cy][x] >= grid[y][x] {
			break
		}
	}

	for cx := x - 1; cx >= 0; cx-- {
		left++
		if grid[y][cx] >= grid[y][x] {
			break
		}
	}

	for cx := x + 1; cx < len(grid); cx++ {
		right++
		if grid[y][cx] >= grid[y][x] {
			break
		}
	}

	return up * down * left * right
}

func B() {
	var grid [][]int

	for line := range utils.ReadDayByLine(8) {
		grid = append(grid, make([]int, len(line)))
		for i, char := range line {
			grid[len(grid)-1][i] = int(char - '0')
		}
	}

	score := make([][]int, len(grid))
	for i := range grid {
		score[i] = make([]int, len(grid[i]))
	}

	for y, row := range grid {
		for x := range row {
			score[y][x] = getTreeScore(x, y, grid)
		}
	}

	result := 0
	for _, row := range score {
		for _, val := range row {
			if val > result {
				result = val
			}
		}
	}

	fmt.Println(result)
}
