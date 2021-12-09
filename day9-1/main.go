package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Solve(lines []string) int {

	num := 0
	heights := make([][]int, len(lines))
	for i, line := range lines {
		heights[i] = make([]int, len(line))
		for j, r := range line {
			n, _ := strconv.Atoi(string(r))
			heights[i][j] = n
		}
	}

	for i, row := range heights {
		for j, point := range row {
			up := math.MaxInt
			if i > 0 {
				up = heights[i-1][j]
			}
			down := math.MaxInt
			if i < len(heights)-1 {
				down = heights[i+1][j]
			}
			left := math.MaxInt
			if j > 0 {
				left = heights[i][j-1]
			}
			right := math.MaxInt
			if j < len(row)-1 {
				right = heights[i][j+1]
			}
			if point < up && point < down && point < left && point < right {
				num += point + 1
			}
		}
	}
	return num
}

func ReadInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	fmt.Println(Solve(ReadInput(os.Stdin)))
}
