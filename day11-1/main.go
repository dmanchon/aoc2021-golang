package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 0 1 2 3
// a b c
// x y z
// s r t

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

		Step(x+1, y, state, flashed)
		Step(x+1, y+1, state, flashed)
		Step(x, y+1, state, flashed)
		Step(x-1, y, state, flashed)
		Step(x-1, y-1, state, flashed)
		Step(x, y-1, state, flashed)
		Step(x+1, y-1, state, flashed)
		Step(x-1, y+1, state, flashed)
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
		empty := make(map[int]bool)

		for x := 0; x < len(state); x++ {
			for y := 0; y < len(state[0]); y++ {
				Step(x, y, state, empty)
			}
		}
		for x := 0; x < len(state); x++ {
			for y := 0; y < len(state[0]); y++ {
				if state[x][y] == 0 {
					num++
				}
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
