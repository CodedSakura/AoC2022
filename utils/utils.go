package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadDayByLine(day int) <-chan string {
	file, err := os.Open(fmt.Sprintf("%02d/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	channel := make(chan string)

	go func() {
		for scanner.Scan() {
			channel <- scanner.Text()
		}
		close(channel)
	}()

	return channel
}

func ReadDayByIntPerLine(day int) <-chan int {
	channel := make(chan int)

	go func() {
		for line := range ReadDayByLine(day) {
			num, _ := strconv.Atoi(line)
			channel <- num
		}
		close(channel)
	}()

	return channel
}

func ChanToArr[T any](channel <-chan T) []T {
	var res []T

	for val := range channel {
		res = append(res, val)
	}

	return res
}

func GetMapKeys[T comparable, V any](m map[T]V) []T {
	res := make([]T, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	return res
}
