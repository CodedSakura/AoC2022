package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now())
	A()
	fmt.Println(time.Now())
	B()
	fmt.Println(time.Now())
}

type pos struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dist(a, b pos) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func A() {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	rowOut := make(map[pos]bool)

	for line := range utils.ReadDayByLine(15) {
		nums := re.FindStringSubmatch(line)[1:]
		sx, _ := strconv.Atoi(nums[0])
		sy, _ := strconv.Atoi(nums[1])
		bx, _ := strconv.Atoi(nums[2])
		by, _ := strconv.Atoi(nums[3])
		d := dist(pos{sx, sy}, pos{bx, by})
		dyOut := dist(pos{sx, sy}, pos{sx, 2_000_000})

		if dyOut > d {
			continue
		}

		dN := d - dyOut

		for x := sx - dN; x <= sx+dN; x++ {
			rowOut[pos{x, 2_000_000}] = true
		}

		if by == 2_000_000 {
			rowOut[pos{bx, by}] = false
		}
	}

	result := 0
	for _, b := range rowOut {
		if b {
			result++
		}
	}

	fmt.Println(result)
}

type sensor struct {
	pos
	dist int
}

func B() {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	var sensors []sensor

	for line := range utils.ReadDayByLine(15) {
		nums := re.FindStringSubmatch(line)[1:]
		sx, _ := strconv.Atoi(nums[0])
		sy, _ := strconv.Atoi(nums[1])
		bx, _ := strconv.Atoi(nums[2])
		by, _ := strconv.Atoi(nums[3])
		d := dist(pos{sx, sy}, pos{bx, by})

		sensors = append(sensors, sensor{pos{sx, sy}, d})
	}

	cx, cy := 0, 0
	for {
		var s sensor
		for _, cs := range sensors {
			if dist(cs.pos, pos{cx, cy}) <= cs.dist {
				s = cs
				break
			}
		}

		if (s == sensor{}) {
			break
		}

		dys := dist(s.pos, pos{s.pos.x, cy})

		cx = s.pos.x + (s.dist - dys)

		cx++
		if cx > 4_000_000 {
			cx = 0
			cy++
		}
	}

	fmt.Println(cx*4000000 + cy)
}
