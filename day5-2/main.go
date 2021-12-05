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

func (l Line) Slope() (int, int) {
	x, y := 0, 0
	if l.x1 < l.x2 {
		x = 1
	}
	if l.y1 < l.y2 {
		y = 1
	}
	if l.x1 > l.x2 {
		x = -1
	}
	if l.y1 > l.y2 {
		y = -1
	}

	return x, y
}

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
		xslope, yslope := line.Slope()
		x, y, i := 0, 0, 0
		for {
			if line.x2 == x && line.y2 == y {
				break
			}
			x = line.x1 + (i * xslope)
			y = line.y1 + (i * yslope)
			pos := maxX*y + x
			coords[pos] += 1
			i++
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
