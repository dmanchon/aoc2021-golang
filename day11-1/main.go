package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Step(x, y int, state [][]int, flashed map[int]bool) {
	if y < 0 || x < 0 || y > len(state[0])-1 || x > len(state)-1 {
		return
	}

	if _, ok := flashed[x+y*len(state)]; ok {
		return
	}

	state[x][y]++
	if state[x][y] > 9 {
		state[x][y] = 0
		flashed[x+y*len(state)] = true

		for _, i := range []int{-1, 0, 1} {
			for _, j := range []int{-1, 0, 1} {
				Step(x+i, y+j, state, flashed)
			}
		}
	}

}

func Solve(lines []string) int {
	num := 0

	state := make([][]int, len(lines))
	for i, line := range lines {
		state[i] = make([]int, len(line))
		for j, r := range line {
			n, _ := strconv.Atoi(string(r))
			state[i][j] = n
		}
	}

	for i := 0; i < 100; i++ {
		flashed := make(map[int]bool)

		for x := 0; x < len(state); x++ {
			for y := 0; y < len(state[0]); y++ {
				Step(x, y, state, flashed)
			}
		}
		num += len(flashed)
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
