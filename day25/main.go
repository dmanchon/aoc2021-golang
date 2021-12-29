package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

func MoveEast(g [][]byte) bool {
	grid := make([][]byte, len(g))
	for i := range g {
		grid[i] = make([]byte, len(g[i]))
		copy(grid[i], g[i])
	}

	modified := false
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			cell := g[i][j]

			if cell != '>' {
				continue
			}
			jp := (j + 1) % len(g[i])
			candidate := g[i][jp]
			if candidate != '.' {
				continue
			}

			grid[i][j] = '.'
			grid[i][jp] = '>'
			modified = true
		}
	}

	for i := range grid {
		for j := range grid[i] {
			g[i][j] = grid[i][j]
		}
	}
	return modified
}

func MoveSouth(g [][]byte) bool {
	grid := make([][]byte, len(g))
	for i := range g {
		grid[i] = make([]byte, len(g[i]))
		copy(grid[i], g[i])
	}

	modified := false
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			cell := g[i][j]

			if cell != 'v' {
				continue
			}
			ip := (i + 1) % len(g)
			candidate := g[ip][j]
			if candidate != '.' {
				continue
			}

			grid[i][j] = '.'
			grid[ip][j] = 'v'
			modified = true
		}
	}

	for i := range grid {
		for j := range grid[i] {
			g[i][j] = grid[i][j]
		}
	}

	return modified
}

func Solve(lines []byte) int {
	grid := bytes.Fields(lines)

	i := 0
	for {
		cond1, cond2 := MoveEast(grid), MoveSouth(grid)
		i++
		if !cond1 && !cond2 {
			return i
		}
	}

}

func ReadInput(file *os.File) []byte {
	reader := bufio.NewReader(file)
	lines, _ := reader.ReadBytes(0)
	return lines
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("- Elapsed Time: ", time.Since(start))
	}()
	fmt.Println("- Solution: ", Solve(ReadInput(os.Stdin)))
}
