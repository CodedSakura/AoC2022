package main

import (
	"AoC2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

func neighborCount(grid [][][]bool, pos [3]int) (out int) {
	for _, n := range getNeighbors(grid, pos) {
		if grid[n[2]][n[1]][n[0]] {
			out++
		}
	}
	return out
}

func getNeighbors[T any](grid [][][]T, pos [3]int) (out [][3]int) {
	if pos[2] > 0 {
		out = append(out, [3]int{pos[0], pos[1], pos[2] - 1})
	}
	if pos[1] > 0 {
		out = append(out, [3]int{pos[0], pos[1] - 1, pos[2]})
	}
	if pos[0] > 0 {
		out = append(out, [3]int{pos[0] - 1, pos[1], pos[2]})
	}
	if pos[2] < len(grid)-1 {
		out = append(out, [3]int{pos[0], pos[1], pos[2] + 1})
	}
	if pos[1] < len(grid[pos[2]])-1 {
		out = append(out, [3]int{pos[0], pos[1] + 1, pos[2]})
	}
	if pos[0] < len(grid[pos[2]][pos[1]])-1 {
		out = append(out, [3]int{pos[0] + 1, pos[1], pos[2]})
	}
	return out
}

func processOutside(isOutside *[][][]int, pos [3]int) {
	visited := make(map[[3]int]bool)
	visited[pos] = true
	queue := [][3]int{pos}
	var v [3]int
	isOut := false
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		if (*isOutside)[v[2]][v[1]][v[0]] == 0 {
			n := getNeighbors(*isOutside, v)
			if len(n) != 6 {
				isOut = true
			}
			for _, w := range n {
				if !visited[w] {
					visited[w] = true
					queue = append(queue, w)
				}
			}
		}
	}
	for c := range visited {
		if (*isOutside)[c[2]][c[1]][c[0]] != 0 {
			continue
		}
		if isOut {
			(*isOutside)[c[2]][c[1]][c[0]] = 3
		} else {
			(*isOutside)[c[2]][c[1]][c[0]] = 1
		}
	}
}

func A() {
	var coordinates [][3]int
	maxCoordinate := [3]int{math.MinInt, math.MinInt, math.MinInt}

	for line := range utils.ReadDayByLine(18) {
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		if x > maxCoordinate[0] {
			maxCoordinate[0] = x
		}
		if y > maxCoordinate[1] {
			maxCoordinate[1] = y
		}
		if z > maxCoordinate[2] {
			maxCoordinate[2] = z
		}
		coordinates = append(coordinates, [3]int{x, y, z})
	}
	grid := make([][][]bool, maxCoordinate[2]+1)
	for z := range grid {
		grid[z] = make([][]bool, maxCoordinate[1]+1)
		for y := range grid[z] {
			grid[z][y] = make([]bool, maxCoordinate[0]+1)
		}
	}
	for _, c := range coordinates {
		grid[c[2]][c[1]][c[0]] = true
	}

	surface := 0
	for _, c := range coordinates {
		surface += 6 - neighborCount(grid, c)
	}

	fmt.Println(surface)
}

func B() {
	var coordinates [][3]int
	maxCoordinate := [3]int{math.MinInt, math.MinInt, math.MinInt}

	for line := range utils.ReadDayByLine(18) {
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		if x > maxCoordinate[0] {
			maxCoordinate[0] = x
		}
		if y > maxCoordinate[1] {
			maxCoordinate[1] = y
		}
		if z > maxCoordinate[2] {
			maxCoordinate[2] = z
		}
		coordinates = append(coordinates, [3]int{x, y, z})
	}
	grid := make([][][]bool, maxCoordinate[2]+1)
	for z := range grid {
		grid[z] = make([][]bool, maxCoordinate[1]+1)
		for y := range grid[z] {
			grid[z][y] = make([]bool, maxCoordinate[0]+1)
		}
	}
	for _, c := range coordinates {
		grid[c[2]][c[1]][c[0]] = true
	}

	totalSurface := 0
	for _, c := range coordinates {
		totalSurface += 6 - neighborCount(grid, c)
	}

	// 0 - unknown, 1 - no, 2 - wall, 3 - yes
	isOutside := make([][][]int, maxCoordinate[2]+1)
	for z := range isOutside {
		isOutside[z] = make([][]int, maxCoordinate[1]+1)
		for y := range isOutside[z] {
			isOutside[z][y] = make([]int, maxCoordinate[0]+1)
			for x := range isOutside[z][y] {
				if grid[z][y][x] {
					isOutside[z][y][x] = 2
				}
			}
		}
	}
	for z, zz := range isOutside {
		for y, yy := range zz {
			for x := range yy {
				if isOutside[z][y][x] != 0 {
					continue
				}
				processOutside(&isOutside, [3]int{x, y, z})
			}
		}
	}

	inside := make([][][]bool, maxCoordinate[2]+1)
	for z := range inside {
		inside[z] = make([][]bool, maxCoordinate[1]+1)
		for y := range inside[z] {
			inside[z][y] = make([]bool, maxCoordinate[0]+1)
			for x := range inside[z][y] {
				inside[z][y][x] = isOutside[z][y][x] < 2
			}
		}
	}

	internalSurface := 0
	for _, c := range coordinates {
		internalSurface += neighborCount(inside, c)
	}

	fmt.Println(totalSurface - internalSurface)
}
