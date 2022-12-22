package main

import (
	"AoC2022/utils"
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	A()
	B()
}

const (
	Right = 0
	Down  = 1
	Left  = 2
	Up    = 3
)

func printPos(board [][]byte, posX, posY, rRot int) {
	for y, row := range board {
		for x, b := range row {
			if x == posX && y == posY {
				switch rRot {
				case Right:
					fmt.Print(">")
					break
				case Down:
					fmt.Print("v")
					break
				case Left:
					fmt.Print("<")
					break
				case Up:
					fmt.Print("^")
					break
				}
				continue
			}
			switch b {
			case 0:
				fmt.Print(" ")
				break
			case 1:
				fmt.Print(".")
				break
			case 2:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func A() {
	var board [][]byte
	var instructions string
	readingBoard := true
	maxBoardWidth := 0

	for line := range utils.ReadDayByLine(22) {
		if line == "" {
			readingBoard = false
			continue
		}

		if readingBoard {
			if len(line) > maxBoardWidth {
				maxBoardWidth = len(line)
			}
			board = append(board, utils.Map([]rune(line), func(c rune) byte {
				switch c {
				case ' ':
					return 0
				case '.':
					return 1
				case '#':
					return 2
				}
				panic(c)
			}))
		} else {
			instructions = line
		}
	}

	for y, row := range board {
		if len(row) < maxBoardWidth {
			board[y] = append(board[y], make([]byte, maxBoardWidth-len(row))...)
		}
	}

	posX := bytes.IndexByte(board[0], 1)
	posY := 0
	rot := Right

	var move func(x, y int)
	move = func(x, y int) {
		for x > 0 {
			posX++
			x--
			if posX >= len(board[posY]) || board[posY][posX] == 0 {
				posX--
				for posX > 0 && board[posY][posX] != 0 {
					posX--
				}
				if board[posY][posX] == 0 {
					posX++
				}
			}
			if board[posY][posX] == 2 {
				move(-1, 0)
				return
			}
		}
		for x < 0 {
			posX--
			x++
			if posX < 0 || board[posY][posX] == 0 {
				posX++
				for posX < len(board[posY]) && board[posY][posX] != 0 {
					posX++
				}
				if posX == len(board[posY]) || board[posY][posX] == 0 {
					posX--
				}
			}
			if board[posY][posX] == 2 {
				move(1, 0)
				return
			}
		}
		for y > 0 {
			posY++
			y--
			if posY >= len(board) || board[posY][posX] == 0 {
				posY--
				for posY > 0 && board[posY][posX] != 0 {
					posY--
				}
				if board[posY][posX] == 0 {
					posY++
				}
			}
			if board[posY][posX] == 2 {
				move(0, -1)
				return
			}
		}
		for y < 0 {
			posY--
			y++
			if posY < 0 || board[posY][posX] == 0 {
				posY++
				for posY < len(board) && board[posY][posX] != 0 {
					posY++
				}
				if posY == len(board) || board[posY][posX] == 0 {
					posY--
				}
			}
			if board[posY][posX] == 2 {
				move(0, 1)
				return
			}
		}
	}

	re := regexp.MustCompile(`\d+|[LR]`)

	for _, i := range re.FindAllStringSubmatch(instructions, -1) {
		instr := i[0]
		if instr == "L" {
			rot = (rot - 1 + 4) % 4
		} else if instr == "R" {
			rot = (rot + 1) % 4
		} else {
			n := utils.ParseInt(instr)
			switch rot {
			case Right:
				move(n, 0)
				break
			case Down:
				move(0, n)
				break
			case Left:
				move(-n, 0)
				break
			case Up:
				move(0, -n)
				break
			}
		}
		//printPos(board, posX, posY)
	}

	fmt.Println((posY+1)*1000 + (posX+1)*4 + rot)
}

func B() {
	var board [][]byte
	var instructions string
	readingBoard := true
	maxBoardWidth := 0

	for line := range utils.ReadDayByLine(22) {
		if line == "" {
			readingBoard = false
			continue
		}

		if readingBoard {
			if len(line) > maxBoardWidth {
				maxBoardWidth = len(line)
			}
			board = append(board, utils.Map([]rune(line), func(c rune) byte {
				switch c {
				case ' ':
					return 0
				case '.':
					return 1
				case '#':
					return 2
				}
				panic(c)
			}))
		} else {
			instructions = line
		}
	}

	for y, row := range board {
		if len(row) < maxBoardWidth {
			board[y] = append(board[y], make([]byte, maxBoardWidth-len(row))...)
		}
	}

	posX := bytes.IndexByte(board[0], 1)
	posY := 0
	rot := Right

	//wraps := [][]int{
	//	/*d x  y  d  x  y  >  <*/
	//	{0, 2, 0, 0, 0, 1, 2, 2},
	//	{1, 3, 0, 1, 4, 2, 2, 2},
	//	{1, 3, 1, 0, 3, 2, 1, 3},
	//	{1, 2, 0, 0, 1, 1, 3, 1},
	//	{1, 0, 1, 0, 3, 3, 1, 3},
	//	{0, 0, 2, 0, 2, 3, 2, 2},
	//	{0, 1, 2, 1, 2, 2, 3, 1},
	//}
	//const size = 4

	wraps := [][]int{
		/*d x  y  d  x  y  >  <*/
		{1, 1, 0, 1, 0, 2, 2, 2},
		{0, 1, 0, 1, 0, 3, 1, 3},
		{0, 2, 0, 0, 0, 4, 0, 0},
		{1, 3, 0, 1, 2, 2, 2, 2},
		{0, 2, 1, 1, 2, 1, 1, 3},
		{1, 1, 1, 0, 0, 2, 3, 1},
		{0, 1, 3, 1, 1, 3, 1, 3},
	}
	const size = 50

	wrap := func() {
		var d, x, y int
		if rot == Left || rot == Right {
			d = 1
			x = (posX + 1) / size
			y = (posY - (posY % size)) / size
		} else {
			d = 0
			x = (posX - (posX % size)) / size
			y = (posY + 1) / size
		}
		var rd, rx, ry, rRot int
		rRot = -1
		for _, wrap := range wraps {
			if wrap[0] == d && wrap[1] == x && wrap[2] == y {
				rd = wrap[3]
				rx = wrap[4]
				ry = wrap[5]
				rRot = wrap[6]
			} else if wrap[3] == d && wrap[4] == x && wrap[5] == y {
				rd = wrap[0]
				rx = wrap[1]
				ry = wrap[2]
				rRot = wrap[7]
			}
		}

		//fmt.Println(posX, posY)
		//printPos(board, posX, posY, rot)
		//fmt.Println(d, x, y)
		//fmt.Println(rd, rx, ry, rRot)
		if rRot == -1 {
			fmt.Println(posX, posY)
			printPos(board, posX, posY, rot)
			fmt.Println(d, x, y)
			fmt.Println(rd, rx, ry, rRot)
			panic(rRot)
		}

		rot = (rot + rRot) % 4
		if d == rd {
			if d == 0 {
				if rRot == 2 {
					posX = (size - (posX % size) - 1) + size*rx
					posY = size * ry
				} else {
					posX = (posX % size) + size*rx
					posY = size * ry
				}
			} else {
				if rRot == 2 {
					posX = size * rx
					posY = (size - (posY % size) - 1) + size*ry
				} else {
					posX = size * rx
					posY = (posY % size) + size*ry
				}
			}
		} else {
			if d == 0 {
				if rRot == 1 { // clockwise
					posY = (posX % size) + size*ry
					posX = size * rx
				} else {
					posY = (size - (posX % size) - 1) + size*ry
					posX = size * rx
				}
			} else {
				if rRot == 3 {
					posX = (posY % size) + size*rx
					posY = size * ry
				} else {
					posX = (size - (posY % size) - 1) + size*rx
					posY = size * ry
				}
			}
		}
		if rot == Up {
			posY--
		}
		if rot == Left {
			posX--
		}
		//fmt.Println(posX, posY)
		//printPos(board, posX, posY, rot)
	}

	var move func(x, y int)

	moveRot := func(n int) {
		switch rot {
		case Right:
			move(n, 0)
			break
		case Down:
			move(0, n)
			break
		case Left:
			move(-n, 0)
			break
		case Up:
			move(0, -n)
			break
		}
	}

	reverseIfInWall := func() bool {
		if board[posY][posX] == 2 {
			//fmt.Println("reversing")
			rot = (rot + 2) % 4
			moveRot(1)
			rot = (rot + 2) % 4
			return true
		}
		return false
	}

	move = func(x, y int) {
		for x > 0 {
			posX++
			x--
			if posX >= len(board[posY]) || board[posY][posX] == 0 {
				wrap()
				if reverseIfInWall() {
					return
				}
				moveRot(x)
				return
			}

			if reverseIfInWall() {
				return
			}
		}

		for x < 0 {
			posX--
			x++
			if posX < 0 || board[posY][posX] == 0 {
				wrap()
				if reverseIfInWall() {
					return
				}
				moveRot(-x)
				return
			}

			if reverseIfInWall() {
				return
			}
		}

		for y > 0 {
			posY++
			y--
			if posY >= len(board) || board[posY][posX] == 0 {
				wrap()
				if reverseIfInWall() {
					return
				}
				moveRot(y)
				return
			}

			if reverseIfInWall() {
				return
			}
		}

		for y < 0 {
			posY--
			y++
			if posY < 0 || board[posY][posX] == 0 {
				wrap()
				if reverseIfInWall() {
					return
				}
				moveRot(-y)
				return
			}

			if reverseIfInWall() {
				return
			}
		}
		//printPos(board, posX, posY, rot)
	}

	re := regexp.MustCompile(`\d+|[LR]`)

	//printPos(board, posX, posY, rot)
	for _, i := range re.FindAllStringSubmatch(instructions, -1) {
		instr := i[0]
		//fmt.Printf("[%d] %s\n", ii, instr)
		//if ii > 991 {
		//	printPos(board, posX, posY, rot)
		//}
		if instr == "L" {
			rot = (rot - 1 + 4) % 4
		} else if instr == "R" {
			rot = (rot + 1) % 4
		} else {
			n := utils.ParseInt(instr)
			moveRot(n)
		}
		//printPos(board, posX, posY, rot)
	}

	fmt.Println((posY+1)*1000 + (posX+1)*4 + rot)
}
