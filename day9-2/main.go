package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

func Adjacent(x1, y1, x2, y2 int) bool {
	if (math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))) == 1 {
		return true
	}
	return false
}

func Merge(basins map[Point][]Point) (*Point, *Point) {
	for k1, basin1 := range basins {
		for k2, basin2 := range basins {
			if k1 == k2 {
				continue
			}
			for _, p1 := range basin1 {
				for _, p2 := range basin2 {
					if Adjacent(p1.x, p1.y, p2.x, p2.y) {
						return &k1, &k2
					}
				}
			}

		}
	}
	return nil, nil
}

func Solve(lines []string) int {

	heights := make([][]bool, len(lines))
	for i, line := range lines {
		heights[i] = make([]bool, len(line))
		for j, r := range line {
			heights[i][j] = (r == '9')
		}
	}

	// initial basins are individual points
	basins := make(map[Point][]Point)
	for i, row := range heights {
		for j, v := range row {
			if !v {
				basins[Point{x: i, y: j}] = []Point{{i, j}}
			}
		}
	}

	// merge basins
	for {
		p1, p2 := Merge(basins)
		if p1 == nil || p2 == nil {
			break
		}
		basins[*p1] = append(basins[*p1], basins[*p2]...)
		delete(basins, *p2)
	}

	// sort desc.
	sizes := make([]int, 0, len(basins))
	for _, basin := range basins {
		sizes = append(sizes, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0] * sizes[1] * sizes[2]
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
