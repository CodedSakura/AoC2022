package main

import (
	"AoC2022/utils"
	"fmt"
	"math"
)

func main() {
	A()
	B()
}

type pos struct {
	x, y int
}

func neighbors(where, max pos) []pos {
	var out []pos

	if where.x > 0 {
		out = append(out, pos{where.x - 1, where.y})
	}
	if where.x < max.x-1 {
		out = append(out, pos{where.x + 1, where.y})
	}
	if where.y > 0 {
		out = append(out, pos{where.x, where.y - 1})
	}
	if where.y < max.y-1 {
		out = append(out, pos{where.x, where.y + 1})
	}

	return out
}

func dijkstra(grid, distance [][]int, end, size pos) {
	q := make(map[pos]struct{})
	for y, row := range distance {
		for x := range row {
			q[pos{x, y}] = struct{}{}
		}
	}

	for len(q) > 0 {
		u := pos{-1, -1}
		uDist := math.MaxInt
		for cPos := range q {
			cDist := distance[cPos.y][cPos.x]
			if cDist < uDist {
				u = cPos
				uDist = cDist
			}
		}
		if u == (pos{-1, -1}) {
			//fmt.Println(q)
			return
			//panic(u)
		}
		delete(q, u)

		if u == end {
			return
		}

		for _, n := range neighbors(u, size) {
			if _, isInQ := q[n]; !isInQ {
				continue
			}
			uHeight, nHeight := grid[u.y][u.x], grid[n.y][n.x]
			if nHeight > uHeight+1 {
				continue
			}
			alt := uDist + 1
			if alt < distance[n.y][n.x] {
				distance[n.y][n.x] = alt
			}
		}
	}
}

func A() {
	var grid [][]int
	var start, end pos

	y := 0
	for line := range utils.ReadDayByLine(12) {
		grid = append(grid, make([]int, len(line)))
		for x, c := range line {
			if c == 'S' {
				start = pos{x, y}
				c = 'a'
			} else if c == 'E' {
				end = pos{x, y}
				c = 'z'
			}
			grid[y][x] = int(c - 'a')
		}
		y++
	}

	size := pos{len(grid[0]), len(grid)}

	var distance [][]int
	for y, row := range grid {
		distance = append(distance, make([]int, len(row)))
		for x := 0; x < len(distance[y]); x++ {
			distance[y][x] = math.MaxInt
		}
	}

	distance[start.y][start.x] = 0

	dijkstra(grid, distance, end, size)

	fmt.Println(distance[end.y][end.x])
}

func B() {
	var grid [][]int
	var end pos

	y := 0
	for line := range utils.ReadDayByLine(12) {
		grid = append(grid, make([]int, len(line)))
		for x, c := range line {
			if c == 'S' {
				c = 'a'
			} else if c == 'E' {
				end = pos{x, y}
				c = 'z'
			}
			grid[y][x] = int(c - 'a')
		}
		y++
	}

	size := pos{len(grid[0]), len(grid)}

	var starts []pos
	for y, row := range grid {
		for x, h := range row {
			if h == 0 {
				starts = append(starts, pos{x, y})
			}
		}
	}

	result := math.MaxInt
	for _, start := range starts {
		var distance [][]int
		for y, row := range grid {
			distance = append(distance, make([]int, len(row)))
			for x := 0; x < len(distance[y]); x++ {
				distance[y][x] = math.MaxInt
			}
		}

		distance[start.y][start.x] = 0

		dijkstra(grid, distance, end, size)

		r := distance[end.y][end.x]
		if r < result {
			result = r
		}
	}
	fmt.Println(result)
}
