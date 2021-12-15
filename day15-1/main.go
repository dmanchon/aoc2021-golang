package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Point struct {
	X, Y int
}

func FindPaths(start, end Point, points map[Point][]Point, costFn func(point Point) int) (int, []Point) {

	dist := make(map[Point]int)
	prev := make(map[Point]*Point)
	queue := make([]Point, 0)

	// init
	for p, _ := range points {
		queue = append(queue, p)
		dist[p] = math.MaxInt
		prev[p] = nil
	}
	dist[end] = 0

	// traverse
	for len(queue) > 0 {
		min := math.MaxInt
		var u Point
		var i int
		for j, p := range queue {
			if dist[p] < min {
				min = dist[p]
				u = p
				i = j
			}
		}

		queue = append(queue[:i], queue[i+1:]...)
		for _, edge := range points[u] {
			for _, v := range queue {
				if v == edge {
					alt := dist[u] + costFn(v)
					if alt < dist[v] {
						dist[v] = alt
						prev[v] = &u
					}
					if v == start {
						path := make([]Point, 0)

						var p *Point = &v
						for {
							path = append(path, *p)

							p = prev[*p]
							if p == nil {
								break
							}
						}
						return dist[v], path
					}
				}
			}
		}

	}

	return math.MaxInt, []Point{}
}

func Solve(lines []string) int {
	coords := make([][]int, len(lines))
	points := make(map[Point][]Point)

	for i, line := range lines {
		coords[i] = make([]int, len(line))
		for j, r := range line {
			n, _ := strconv.Atoi(string(r))
			coords[i][j] = n
			point := Point{X: i, Y: j}
			points[point] = []Point{}
		}
	}

	for point, edges := range points {
		if point.Y > 0 {
			edges = append(edges, Point{X: point.X, Y: point.Y - 1})
		}
		if point.X > 0 {
			edges = append(edges, Point{X: point.X - 1, Y: point.Y})
		}
		points[point] = edges
	}

	costFn := func(point Point) int {
		return coords[point.X][point.Y]
	}
	_, path := FindPaths(Point{0, 0}, Point{len(coords[0]) - 1, len(coords) - 1}, points, costFn)

	sum := 0
	for _, p := range path[1:] {
		sum += coords[p.X][p.Y]
	}
	return sum
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
	start := time.Now()
	defer func() {
		fmt.Println("- Elapsed Time: ", time.Now().Sub(start))
	}()
	fmt.Println("- Solution: ", Solve(ReadInput(os.Stdin)))
}
