package main

import (
	"AoC2022/utils"
	"fmt"
)

func main() {
	A()
	B()
}

type node struct {
	nextNode, prevNode *node
	value, index       int
}

func A() {
	strArr := utils.ChanToArr(utils.ReadDayByLine(20))

	l := len(strArr)
	arr := make([]node, l)
	sequence := make([]int, l)
	var zero *node

	for i, s := range strArr {
		a := utils.ParseInt(s)
		arr[i] = node{value: a, index: i}
		sequence[i] = a
		if a == 0 {
			zero = &arr[i]
		}
	}

	for i := range arr {
		arr[i].prevNode = &arr[(i-1+l)%l]
		arr[i].nextNode = &arr[(i+1)%l]
	}

	for index, offset := range sequence {
		cNode := &arr[index]

		if offset == 0 {
			continue
		}

		cNode.prevNode.nextNode = cNode.nextNode
		cNode.nextNode.prevNode = cNode.prevNode

		nNode := cNode
		if offset > 0 {
			for i := 0; i <= offset; i++ {
				nNode = nNode.nextNode
			}
		} else if offset < 0 {
			for i := 0; i > offset; i-- {
				nNode = nNode.prevNode
			}
		}

		cNode.prevNode = nNode.prevNode
		cNode.nextNode = nNode
		nNode.prevNode = cNode
		cNode.prevNode.nextNode = cNode
	}

	cNode := zero
	nArr := make([]int, l)
	nI := 0
	for cNode.nextNode != zero {
		nArr[nI] = cNode.value
		cNode = cNode.nextNode
		nI++
	}
	nArr[nI] = cNode.value

	fmt.Println(
		nArr[1000%l] +
			nArr[2000%l] +
			nArr[3000%l],
	)
}

func B() {
	strArr := utils.ChanToArr(utils.ReadDayByLine(20))

	l := len(strArr)
	arr := make([]node, l)
	sequence := make([]int, l)
	var zero *node

	for i, s := range strArr {
		a := utils.ParseInt(s) * 811589153
		arr[i] = node{value: a, index: i}
		sequence[i] = a
		if a == 0 {
			zero = &arr[i]
		}
	}

	for i := range arr {
		arr[i].prevNode = &arr[(i-1+l)%l]
		arr[i].nextNode = &arr[(i+1)%l]
	}

	for _i := 0; _i < 10; _i++ {
		for index, offset := range sequence {
			cNode := &arr[index]

			if offset == 0 {
				continue
			}

			cNode.prevNode.nextNode = cNode.nextNode
			cNode.nextNode.prevNode = cNode.prevNode

			offset = offset % (l - 1)

			nNode := cNode
			if offset > 0 {
				for i := 0; i <= offset; i++ {
					nNode = nNode.nextNode
				}
			} else if offset < 0 {
				for i := 0; i > offset; i-- {
					nNode = nNode.prevNode
				}
			}

			cNode.prevNode = nNode.prevNode
			cNode.nextNode = nNode
			nNode.prevNode = cNode
			cNode.prevNode.nextNode = cNode
		}
	}

	cNode := zero
	nArr := make([]int, l)
	nI := 0
	for cNode.nextNode != zero {
		nArr[nI] = cNode.value
		cNode = cNode.nextNode
		nI++
	}
	nArr[nI] = cNode.value

	fmt.Println(
		nArr[1000%l] +
			nArr[2000%l] +
			nArr[3000%l],
	)
}
