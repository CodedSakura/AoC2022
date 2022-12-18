package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	A()
	B()
}

type valve struct {
	name         string
	flowRate     int
	plainTunnels []string
	tunnels      map[*valve]int
}

func solveA(aa *valve) int {
	res := 0

	var dfs func(v *valve, time, score int, discovered map[*valve]bool)
	dfs = func(v *valve, time, score int, discovered map[*valve]bool) {
		discovered[v] = true
		for w, d := range v.tunnels {
			if !discovered[w] {
				t := time - d - 1
				if t < 0 {
					if score > res {
						res = score
					}
					continue
				}
				s := score + t*w.flowRate
				if s > res {
					res = s
				}
				//fmt.Println(discovered)
				dCopy := make(map[*valve]bool, len(discovered))
				for k, w := range discovered {
					dCopy[k] = w
				}
				dfs(w, t, s, dCopy)
			}
		}
	}

	discovered := make(map[*valve]bool)
	dfs(aa, 30, 0, discovered)

	return res
}

func solveB(aa *valve) int {
	// me  ->  j b c
	// ele -> d  h  e

	maxRes := 0
	counter := 0

	getDiscoveredKey := func(discovered map[*valve]bool) string {
		var names []string
		for k, d := range discovered {
			if !d {
				continue
			}
			names = append(names, k.name)
		}
		sort.Strings(names)
		return strings.Join(names, ":")
	}

	dp := make(map[string]int)

	var dfs func(v *valve, time, score int, discovered map[*valve]bool, amElephant bool) int
	dfs = func(v *valve, time, score int, discovered map[*valve]bool, amElephant bool) int {
		res := 0
		discovered[v] = true
		counter++
		if counter%1_000_000 == 0 {
			fmt.Println(counter)
		}
		if !amElephant {
			key := getDiscoveredKey(discovered)
			var eRes int
			if dpV, e := dp[key]; e {
				eRes = score + dpV
				if eRes > maxRes {
					maxRes = eRes
				}
			} else {
				dECopy := make(map[*valve]bool, len(discovered))
				for k, w := range discovered {
					dECopy[k] = w
				}
				eRes = dfs(aa, 26, 0, dECopy, true)
				dp[key] = eRes
				eRes += score
				//eRes := dfs(aa, 26, score, dECopy, true)
			}
			if eRes > maxRes {
				maxRes = eRes
			}
		}
		for w, d := range v.tunnels {
			if !discovered[w] {
				t := time - d - 1
				if t < 0 {
					if score > res {
						res = score
					}
					continue
				}
				s := score + t*w.flowRate
				if s > res {
					res = s
				}
				//fmt.Println(discovered)
				dCopy := make(map[*valve]bool, len(discovered))
				for k, w := range discovered {
					dCopy[k] = w
				}
				s2 := dfs(w, t, s, dCopy, amElephant)
				if s2 > res {
					res = s2
				}
			}
		}
		return res
	}

	discovered := make(map[*valve]bool)
	dfs(aa, 26, 0, discovered, false)

	fmt.Println(len(dp))
	return maxRes
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
	for k := range visited {
		for i := range visited {
			for j := range visited {
				if _, e := i.tunnels[k]; !e {
					continue
				}
				if _, e := k.tunnels[j]; !e {
					continue
				}
				if i == j || j == k || i == k {
					continue
				}
				dist := i.tunnels[k] + k.tunnels[j]
				if v, e := i.tunnels[j]; !e || v > dist {
					i.tunnels[j] = dist
				}
			}
		}
	}
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

	fmt.Println(solveA(aa))
}

func B() {
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

	fmt.Println(solveB(aa))
}
