package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculation(done chan int64, matrix [][]rune, inverse bool) {
	i := 0
	for len(matrix) > 1 {
		tmp := make([][]rune, 0)

		ones := make([]int, 0)
		zeroes := make([]int, 0)

		for j, row := range matrix {
			switch row[i] {
			case '0':
				zeroes = append(zeroes, j)

			case '1':
				ones = append(ones, j)
			}
		}

		cond := len(ones) >= len(zeroes)
		if inverse {
			// only different between the two cases
			cond = !cond
		}

		if cond {
			for _, idx := range ones {
				tmp = append(tmp, matrix[idx])
			}
		} else {
			for _, idx := range zeroes {
				tmp = append(tmp, matrix[idx])
			}
		}
		matrix = tmp
		i = i + 1
	}

	result, _ := strconv.ParseInt(string(matrix[0]), 2, 64)
	done <- result
}

func solve(lines []string) int64 {
	size := len(lines)
	width := len(lines[0])
	matrix := make([][]rune, size)

	for i, line := range lines {
		matrix[i] = make([]rune, width)
		for j, cell := range line {
			matrix[i][j] = cell
		}
	}

	co2 := make(chan int64, 1)
	o2 := make(chan int64, 1)

	go calculation(o2, matrix, false)
	go calculation(co2, matrix, true)

	return <-o2 * <-co2
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
