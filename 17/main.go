package main

import (
	"AoC2022/utils"
	"fmt"
	"time"
)

func main() {
	A()
	B()
}

type rock struct {
	data []byte
	posY int // bottom of the rock
}

// rocks are inverted
var rockMap = [5][]byte{
	{
		0b0011110,
	},
	{
		0b0001000,
		0b0011100,
		0b0001000,
	},
	{
		0b0011100,
		0b0000100,
		0b0000100,
	},
	{
		0b0010000,
		0b0010000,
		0b0010000,
		0b0010000,
	},
	{
		0b0011000,
		0b0011000,
	},
}

func dropRock(tower *[]byte, index int) (newRock rock) {
	index %= 5
	topPoint := len(*tower) - 1

	// find top point
	for topPoint >= 0 && (*tower)[topPoint] == 0 {
		topPoint--
	}

	// make space for new piece
	newRockSize := len(rockMap[index])
	topPoint += 3 + newRockSize
	for len(*tower) <= topPoint {
		*tower = append(*tower, 0)
	}

	// add piece
	newRock.posY = topPoint - newRockSize + 1
	newRock.data = make([]byte, newRockSize)
	copy(newRock.data, rockMap[index])
	return
}

func settleRock(tower *[]byte, fallingRock rock) {
	for i, r := range fallingRock.data {
		(*tower)[fallingRock.posY+i] |= r
	}
}

//func imposeRock(tower []int, fallingRock rock) []int {
//	out := make([]int, len(tower))
//	copy(out, tower)
//	settleRock(&out, fallingRock)
//	return out
//}

func moveRock(fallingRock *rock, tower *[]byte, jets []int, jetIndex int) (didSettle bool) {
	jetIndex %= len(jets)
	currentJet := jets[jetIndex]
	canMove := true
	// jets
	if currentJet == 1 { // >
		for i, r := range fallingRock.data {
			// will hit wall
			if r&1 == 1 {
				canMove = false
				break
			}
			// will hit tower
			if r>>1&(*tower)[fallingRock.posY+i] != 0 {
				canMove = false
				break
			}
		}
		if canMove {
			for i := range fallingRock.data {
				fallingRock.data[i] >>= 1
			}
		}
	} else { // <
		for i, r := range fallingRock.data {
			// will hit wall
			if r&0b1000000 != 0 {
				canMove = false
				break
			}
			// will hit tower
			if r<<1&(*tower)[fallingRock.posY+i] != 0 {
				canMove = false
				break
			}
		}
		if canMove {
			for i := range fallingRock.data {
				fallingRock.data[i] <<= 1
			}
		}
	}

	// gravity
	if fallingRock.posY == 0 {
		didSettle = true
		settleRock(tower, *fallingRock)
		return
	}
	canFall := true
	for i, r := range fallingRock.data {
		if (*tower)[fallingRock.posY+i-1]&r != 0 {
			canFall = false
			break
		}
	}
	if canFall {
		fallingRock.posY--
	} else {
		didSettle = true
		settleRock(tower, *fallingRock)
	}

	return
}

//func printTower(tower []int) {
//	for i := range tower {
//		s := fmt.Sprintf("%07b", tower[len(tower)-1-i])
//		s = strings.ReplaceAll(s, "0", ".")
//		s = strings.ReplaceAll(s, "1", "#")
//		fmt.Println(s)
//	}
//	fmt.Println()
//}

func A() {
	line := utils.ChanToArr(utils.ReadDayByLine(17))[0]
	jets := make([]int, len(line))
	for i, s := range line {
		if s == '<' {
			jets[i] = -1
		} else {
			jets[i] = 1
		}
	}

	pieceIndex := 0
	jetIndex := 0
	var tower []byte
	for i := 0; i < 2022; i++ {
		fRock := dropRock(&tower, pieceIndex)
		for !moveRock(&fRock, &tower, jets, jetIndex) {
			jetIndex++
		}
		jetIndex++
		pieceIndex++
	}

	top := len(tower) - 1
	for tower[top] == 0 {
		top--
	}
	top++
	fmt.Println(top)
}

func B() {
	line := utils.ChanToArr(utils.ReadDayByLine(17))[0]
	jets := make([]int, len(line))
	for i, s := range line {
		if s == '<' {
			jets[i] = -1
		} else {
			jets[i] = 1
		}
	}

	pieceIndex := 0
	jetIndex := 0
	//tower := make([]byte, 1024*1024)
	var tower []byte
	//towerExtraHeight := 0
	for i := 0; i < 1_000_000_000_000; i++ {
		fRock := dropRock(&tower, pieceIndex)

		if fRock.posY > 1024*1024-10 {
			//	towerExtraHeight += 1024 * 512
			tower = tower[1024*512:]
			fRock.posY -= 1024 * 512
			//	tower = append(tower, make([]byte, 1024*512)...)
		}

		for !moveRock(&fRock, &tower, jets, jetIndex) {
			jetIndex++
		}
		jetIndex++
		pieceIndex++

		if i%50_000_000 == 0 {
			fmt.Printf("[%s] %d / %d (%.03f%%)\n", time.Now().Format("15:04:05.000"), i, 1_000_000_000_000, float64(i)/float64(10_000_000_000))
		}
	}

	top := len(tower) - 1
	for tower[top] == 0 {
		top--
	}
	top++
	fmt.Println(top)
}
