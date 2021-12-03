package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var mapping = map[rune]float64{
	'0': 0.0,
	'1': 1.0,
}

func solve(lines []string) int {
	size := len(lines)
	width := len(lines[0])
	matrix := make([][]rune, size)
	var tmp [][]rune

	for i, line := range lines {
		matrix[i] = make([]rune, width)
		for j, cell := range line {
			matrix[i][j] = cell
		}
	}

	store := matrix

	for i := 0; i < width; i++ {
		if len(matrix) == 1 {
			break
		}

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

		tmp = make([][]rune, 0)
		if len(ones) >= len(zeroes) {
			for _, idx := range ones {
				tmp = append(tmp, matrix[idx])
			}
		} else {
			for _, idx := range zeroes {
				tmp = append(tmp, matrix[idx])
			}
		}
		matrix = tmp
	}
	o2, _ := strconv.ParseInt(string(matrix[0]), 2, 64)

	matrix = store
	for i := 0; i < width; i++ {
		if len(matrix) == 1 {
			break
		}
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

		tmp = make([][]rune, 0)
		if len(zeroes) <= len(ones) {
			for _, idx := range zeroes {
				tmp = append(tmp, matrix[idx])
			}
		} else {
			for _, idx := range ones {
				tmp = append(tmp, matrix[idx])
			}
		}
		matrix = tmp
	}
	co2, _ := strconv.ParseInt(string(matrix[0]), 2, 64)

	return int(co2) * int(o2)
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
