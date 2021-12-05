package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Lines []Line

func (l Lines) Max() (int, int) {
	var x, y int
	for _, coord := range l {
		if x < coord.x1 {
			x = coord.x1
		}
		if x < coord.x2 {
			x = coord.x2
		}

		if y < coord.y1 {
			y = coord.x1
		}
		if y < coord.y2 {
			y = coord.y2
		}

	}
	return x, y
}

func solve(ls []string) int {
	lines := make(Lines, 0, len(ls))

	for _, l := range ls {
		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		for _, match := range re.FindAllStringSubmatch(l, -1) {
			x1, _ := strconv.Atoi(match[1])
			y1, _ := strconv.Atoi(match[2])
			x2, _ := strconv.Atoi(match[3])
			y2, _ := strconv.Atoi(match[4])

			line := Line{
				x1: x1,
				y1: y1,
				x2: x2,
				y2: y2,
			}
			lines = append(lines, line)
		}
	}

	maxX, maxY := lines.Max()
	coords := make([]int, maxY*maxX+maxX)

	for _, line := range lines {
		if line.x1 == line.x2 {
			// vertical
			a, b := line.y2, line.y1
			if line.y2 > line.y1 {
				a, b = line.y1, line.y2
			}
			for i := a; i <= b; i++ {
				pos := maxX*i + line.x2
				coords[pos] += 1
			}
		}

		if line.y1 == line.y2 {
			//horizontal
			a, b := line.x2, line.x1
			if line.x2 > line.x1 {
				a, b = line.x1, line.x2
			}
			for i := a; i <= b; i++ {
				pos := maxX*line.y2 + i
				coords[pos] += 1
			}
		}

	}

	sum := 0
	for _, c := range coords {
		if c > 1 {
			sum++
		}
	}

	return sum
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
