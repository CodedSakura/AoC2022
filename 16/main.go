package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	A()
	// Template file
}

type valve struct {
	name         string
	flowRate     int
	plainTunnels []string
	tunnels      map[*valve]int
	isOpen       bool
}

func findNextBestValve(start *valve, timeLeft int) (end *valve, distance int) {
	parent := make(map[*valve]*valve)
	parent[start] = nil
	queue := []*valve{start}
	for len(queue) > 0 {
		var s *valve
		s, queue = queue[0], queue[1:]

		for v := range s.tunnels {
			if _, visited := parent[v]; !visited {
				queue = append(queue, v)
				parent[v] = s
			}
		}
	}

	maxScore := -1
	var maxValve *valve
	maxDist := 0
	for v := range parent {
		if v.isOpen || v.flowRate == 0 {
			continue
		}

		dist := 0
		c := v
		for c != nil {
			c = parent[c]
			dist++
		}

		score := (timeLeft - dist - 1) * v.flowRate
		fmt.Println(v, score, dist)
		if score > maxScore {
			maxScore = score
			maxValve = v
			maxDist = dist
		}
	}
	return maxValve, maxDist
}

func printGraph(aa *valve) {
	visited := make(map[*valve]bool)
	visited[aa] = true
	queue := []*valve{aa}
	for len(queue) > 0 {
		var s *valve
		s, queue = queue[0], queue[1:]
		if s.flowRate == 0 {
			fmt.Printf("%s [color=cyan,style=filled]\n", s.name)
		}

		for v, d := range s.tunnels {
			fmt.Printf("%s -> %s [label=%d]\n", s.name, v.name, d)
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}
}

func collapse(aa *valve) {
	visited := make(map[*valve]bool)
	visited[aa] = true
	queue := []*valve{aa}
	var s *valve
	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]
		if s.flowRate == 0 && s != aa {
			if len(s.tunnels) != 2 {
				panic(s.tunnels)
			}

			var a, b *valve
			for v := range s.tunnels {
				if a == nil {
					a = v
				} else {
					b = v
				}
			}

			a.tunnels[b] = s.tunnels[a] + s.tunnels[b]
			b.tunnels[a] = s.tunnels[a] + s.tunnels[b]
			delete(a.tunnels, s)
			delete(b.tunnels, s)
		}

		for v := range s.tunnels {
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}

	// collapse AA
	for a, la := range aa.tunnels {
		for b, lb := range aa.tunnels {
			if a == b {
				continue
			}

			a.tunnels[b] = la + lb
			b.tunnels[a] = la + lb
			delete(a.tunnels, aa)
			delete(b.tunnels, aa)
		}
	}

	// Floyd-Warshall
	// TODO
}

func A() {
	re := regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

	valveMap := make(map[string]*valve)

	for line := range utils.ReadDayByLine(16) {
		res := re.FindStringSubmatch(line)[1:]
		rate, _ := strconv.Atoi(res[1])
		valveMap[res[0]] = &valve{
			name:         res[0],
			flowRate:     rate,
			plainTunnels: strings.Split(res[2], ", "),
		}
	}

	// setup tunnels
	for _, v := range valveMap {
		v.tunnels = make(map[*valve]int, len(v.plainTunnels))
		for _, s := range v.plainTunnels {
			v2 := valveMap[s]
			v.tunnels[v2] = 1
		}
		v.plainTunnels = []string{}
	}
	aa := valveMap["AA"]

	collapse(aa)

	printGraph(aa)

	fmt.Println(findNextBestValve(aa, 30))

	//fmt.Println(aa.tunnelsTo[2])
}
