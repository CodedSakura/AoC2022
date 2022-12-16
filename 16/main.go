package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
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
}

func solveA(aa *valve) int {
	res := 0
	nodes := utils.GetMapKeys(aa.tunnels)
	count := 0

	nodes = append([]*valve{aa}, nodes...)
	mat := make([][]int, len(nodes))
	for a, ka := range nodes {
		mat[a] = make([]int, len(nodes))
		for b, kb := range nodes {
			if ka == kb {
				mat[a][b] = 0
			} else {
				mat[a][b] = ka.tunnels[kb]
			}
		}
		//fmt.Println(mat[a])
	}
	//fmt.Println()
	rates := make([]int, len(nodes))
	keys := make([]int, len(nodes))
	for i, v := range nodes {
		rates[i] = v.flowRate
		keys[i] = i
	}
	//fmt.Println(rates)
	//fmt.Println(keys)
	//fmt.Println()

	run := func(p []int, lCount *int) {
		var cRes, cTime int
		var cPos, nPos int
		if *lCount%500_000_000 == 0 {
			fmt.Printf(
				"[%s]: %d / %d (%f%%)\n",
				time.Now().Format("15:04:05.000"), *lCount, 1307674368000/15, float64(*lCount)/(13076743680/15),
			)
		}
		//if count == 1_000_000_000 {
		//	panic("(:")
		//}
		*lCount++

		cRes = 0
		cTime = 30
		cPos = 0
		for _, nPos = range p {
			cTime -= mat[cPos][nPos] + 1
			if cTime < 0 {
				break
			}
			//fmt.Printf("%d,", cTime*rates[nPos])
			cRes += cTime * rates[nPos]
			cPos = nPos
		}
		//fmt.Println(p)

		if cRes > res {
			res = cRes
		}
	}

	wg := sync.WaitGroup{}
	keys = keys[1:]
	permuteFixedLast := func(keys []int) {
		defer wg.Done()

		n := len(keys) - 1
		c := make([]int, n)

		i := 1
		lCount := 0
		run(keys, &lCount)

		for i < n {
			if c[i] < i {
				if i%2 == 0 {
					keys[0], keys[i] = keys[i], keys[0]
				} else {
					keys[c[i]], keys[i] = keys[i], keys[c[i]]
				}
				run(keys, &lCount)
				c[i]++
				i = 1
			} else {
				c[i] = 0
				i++
			}
		}
	}

	for i := range keys {
		cKeys := make([]int, len(keys))
		copy(cKeys, keys)
		cKeys[i], cKeys[len(cKeys)-1] = cKeys[len(cKeys)-1], cKeys[i]
		wg.Add(1)
		go permuteFixedLast(cKeys)
	}
	wg.Wait()
	fmt.Println(count)

	return res
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

	//f, err := os.Create("prof00.out")
	//if err != nil {
	//	log.Fatal("could not create CPU profile: ", err)
	//}
	//defer f.Close()
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal("could not start CPU profile: ", err)
	//}
	//defer pprof.StopCPUProfile()

	// (AA) DD BB JJ HH EE CC
	//    0 20 13 21 22  3  2

	collapse(aa)

	fmt.Println(solveA(aa))
}
