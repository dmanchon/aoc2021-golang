package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Step(x, y int, state [][]int, flashed map[int]bool) {
	if y < 0 || x < 0 || x > len(state)-1 || y > len(state[0])-1 {
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
	state := make([][]int, len(lines))
	for i, line := range lines {
		state[i] = make([]int, len(line))
		for j, r := range line {
			n, _ := strconv.Atoi(string(r))
			state[i][j] = n
		}
	}

	i := 0
	for {
		i++
		empty := make(map[int]bool)

		for x := 0; x < len(state); x++ {
			for y := 0; y < len(state[0]); y++ {
				Step(x, y, state, empty)
			}
		}

		// everything's illuminated?
		sum := len(state) * len(state[0])
		for _, _ = range empty {
			sum--
		}
		if sum == 0 {
			return i
		}
	}

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