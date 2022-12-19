package main

import (
	"AoC2022/utils"
	"fmt"
	"regexp"
	"sync"
)

func main() {
	A()
	B()
}

type blueprint struct {
	index                                     int
	oreRobotOreCost                           int
	clayRobotOreCost                          int
	obsRobotOreCost, obsRobotClayCost         int
	geodeRobotOreCost, geodeRobotObsidianCost int
}

type factory struct {
	oreRobots, clayRobots, obsRobots, geodeRobots int
	ore, clay, obs, geode                         int
}

func processFactoryMinute(fac *factory) {
	fac.ore += fac.oreRobots
	fac.clay += fac.clayRobots
	fac.obs += fac.obsRobots
	fac.geode += fac.geodeRobots
}

func processBlueprint(bp blueprint, time int) int {
	var dfs func(minute int, fac factory) factory
	dfs = func(minute int, fac factory) factory {
		if minute == time {
			return fac
		}
		minute++
		bestFac := fac
		//fmt.Println(minute)

		if fac.ore >= bp.geodeRobotOreCost && fac.obs >= bp.geodeRobotObsidianCost {
			geodeFac := fac
			geodeFac.ore -= bp.geodeRobotOreCost
			geodeFac.obs -= bp.geodeRobotObsidianCost
			processFactoryMinute(&geodeFac)
			geodeFac.geodeRobots += 1
			return dfs(minute, geodeFac)
		}

		if fac.ore >= bp.obsRobotOreCost && fac.clay >= bp.obsRobotClayCost {
			obsFac := fac
			obsFac.ore -= bp.obsRobotOreCost
			obsFac.clay -= bp.obsRobotClayCost
			processFactoryMinute(&obsFac)
			obsFac.obsRobots += 1
			obsFac = dfs(minute, obsFac)
			if obsFac.geode > bestFac.geode {
				bestFac = obsFac
			}
		}

		if fac.clayRobots < bp.obsRobotClayCost && fac.ore >= bp.clayRobotOreCost {
			clayFac := fac
			clayFac.ore -= bp.clayRobotOreCost
			processFactoryMinute(&clayFac)
			clayFac.clayRobots += 1
			clayFac = dfs(minute, clayFac)
			if clayFac.geode > bestFac.geode {
				bestFac = clayFac
			}
		}

		if fac.oreRobots < 5 && fac.ore >= bp.oreRobotOreCost {
			oreFac := fac
			oreFac.ore -= bp.oreRobotOreCost
			processFactoryMinute(&oreFac)
			oreFac.oreRobots += 1
			oreFac = dfs(minute, oreFac)
			if oreFac.geode > bestFac.geode {
				bestFac = oreFac
			}
		}

		if fac.ore < 5 {
			plainFac := fac
			processFactoryMinute(&plainFac)
			plainFac = dfs(minute, plainFac)
			if plainFac.geode > bestFac.geode {
				bestFac = plainFac
			}
		}

		return bestFac
	}

	return dfs(0, factory{oreRobots: 1}).geode
}

func A() {
	re := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	var blueprints []blueprint

	for line := range utils.ReadDayByLine(19) {
		strings := re.FindStringSubmatch(line)[1:]
		blueprints = append(blueprints, blueprint{
			index:                  utils.ParseInt(strings[0]),
			oreRobotOreCost:        utils.ParseInt(strings[1]),
			clayRobotOreCost:       utils.ParseInt(strings[2]),
			obsRobotOreCost:        utils.ParseInt(strings[3]),
			obsRobotClayCost:       utils.ParseInt(strings[4]),
			geodeRobotOreCost:      utils.ParseInt(strings[5]),
			geodeRobotObsidianCost: utils.ParseInt(strings[6]),
		})
	}

	var wg sync.WaitGroup

	result := 0
	for _, bp := range blueprints {
		wg.Add(1)
		go func(bp blueprint) {
			defer wg.Done()
			result += processBlueprint(bp, 24) * bp.index
		}(bp)
	}
	wg.Wait()

	fmt.Println(result)
}

func B() {
	re := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	var blueprints []blueprint

	for line := range utils.ReadDayByLine(19) {
		strings := re.FindStringSubmatch(line)[1:]
		blueprints = append(blueprints, blueprint{
			index:                  utils.ParseInt(strings[0]),
			oreRobotOreCost:        utils.ParseInt(strings[1]),
			clayRobotOreCost:       utils.ParseInt(strings[2]),
			obsRobotOreCost:        utils.ParseInt(strings[3]),
			obsRobotClayCost:       utils.ParseInt(strings[4]),
			geodeRobotOreCost:      utils.ParseInt(strings[5]),
			geodeRobotObsidianCost: utils.ParseInt(strings[6]),
		})
	}

	var wg sync.WaitGroup

	result := 1
	for _, bp := range blueprints[:3] {
		wg.Add(1)
		go func(bp blueprint) {
			defer wg.Done()
			result *= processBlueprint(bp, 32)
		}(bp)
	}
	wg.Wait()

	fmt.Println(result)
}
