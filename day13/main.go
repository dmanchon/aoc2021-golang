package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func Print(x, y int, coords map[Point]struct{}) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if _, ok := coords[Point{X: i, Y: j}]; !ok {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func FoldX(coords map[Point]struct{}, target int) map[Point]struct{} {
	folded := make(map[Point]struct{})
	for p, _ := range coords {
		np := Point{X: p.X, Y: p.Y}
		if p.Y > target {
			np.Y = 2*target - p.Y
		}
		folded[np] = struct{}{}
	}
	return folded
}

func FoldY(coords map[Point]struct{}, target int) map[Point]struct{} {
	folded := make(map[Point]struct{})
	for p, _ := range coords {
		np := Point{X: p.X, Y: p.Y}

		if p.X > target {
			np.X = 2*target - p.X
		}
		folded[np] = struct{}{}
	}
	return folded
}

func Solve(lines []string) int {
	coords := make(map[Point]struct{})
	i := 0
	for _, line := range lines {
		i++
		if line == "" {
			break
		}
		s := strings.Split(line, ",")

		y, _ := strconv.Atoi(s[0])
		x, _ := strconv.Atoi(s[1])

		coords[Point{X: x, Y: y}] = struct{}{}
	}

	for _, line := range lines[i:] {
		s := strings.Split(line, "=")
		target, _ := strconv.Atoi(s[1])

		if s[0][len(s[0])-1] == 'x' {
			coords = FoldX(coords, target)
		} else {
			coords = FoldY(coords, target)
		}
	}
	Print(10, 80, coords)
	return len(coords)
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
