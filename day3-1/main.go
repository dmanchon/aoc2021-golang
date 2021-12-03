package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mapping = map[rune]int{
	'0': 0,
	'1': 1,
}

func solve(lines []string) int {
	size := len(lines)
	width := len(lines[0])
	digits := make([]int, width)
	gamma := make([]rune, width)
	epsilon := make([]rune, width)

	for _, line := range lines {
		for i, digit := range line {
			digits[i] += mapping[digit]
		}

	}

	for i := 0; i < width; i++ {
		if digits[i] > (size / 2) {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	e, _ := strconv.ParseInt(string(epsilon), 2, 64)
	return int(g * e)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	solution := solve(lines)
	fmt.Println(solution)
}
