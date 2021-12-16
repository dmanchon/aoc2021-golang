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

func ReconstructPath(cameFrom map[Point]Point, current Point) []Point {
	path := make([]Point, 0)
	path = append(path, current)
	for {

		next, ok := cameFrom[current]
		if !ok {
			break
		}
		current = next
		path = append(path, current)

	}
	return path
}

func FindPaths(coords [][]int, start, goal Point, points map[Point][]Point, d func(point Point) int) []Point {
	// heuristic function
	h := func(neighbor Point) int {
		//return (goal.X - neighbor.X) ^ 2 + (goal.Y - neighbor.Y) ^ 2
		return 0
	}
	openSet := make(map[Point]struct{})
	openSet[start] = struct{}{}
	cameFrom := make(map[Point]Point)

	gScore := make(map[Point]int)
	for k, _ := range points {
		gScore[k] = math.MaxInt
	}
	gScore[start] = 0

	fScore := make(map[Point]int)
	for k, _ := range points {
		fScore[k] = math.MaxInt
	}
	fScore[start] = h(start)

	for len(openSet) > 0 {

		var min int = math.MaxInt
		var current Point
		for k, _ := range openSet {
			if fScore[k] < min {
				min = fScore[k]
				current = k
			}
		}
		if current == goal {
			return ReconstructPath(cameFrom, current)
		}

		delete(openSet, current)
		for _, neighbor := range points[current] {
			tentative_gScore := gScore[current] + d(neighbor)
			if tentative_gScore < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentative_gScore
				fScore[neighbor] = tentative_gScore + h(neighbor)
				openSet[neighbor] = struct{}{}
			}

		}
		// Uncomment for Live Animation.
		//Print(coords, len(coords), len(coords[0]), ReconstructPath(cameFrom, current), openSet)
	}

	return []Point{}
}

func Print(arr [][]int, x, y int, path []Point, openSet map[Point]struct{}) {
	fmt.Print("\033[H\033[2J")
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			for _, p := range path {
				if p.X == j && p.Y == i {
					fmt.Printf("\033[45m\033[1;33m")
					break
				} else if _, ok := openSet[Point{X: i, Y: j}]; ok {
					fmt.Printf("\033[47m\033[1;32m")
					break
				}
			}
			fmt.Printf("%d", arr[j][i])
			fmt.Printf("\033[0m")
		}
		fmt.Printf("\n")
	}
	time.Sleep(time.Second / 24)
}

func Expand(arr [][]int) [][]int {
	expanded := make([][]int, len(arr)*5)
	for i := range expanded {
		expanded[i] = make([]int, len(arr[0])*5)
		for j := range expanded[i] {
			y := i % len(arr)
			x := j % len(arr[0])
			v := int(j/len(arr[0])) + int(i/len(arr))
			expanded[i][j] = arr[y][x] + v

			if expanded[i][j] > 9 {
				expanded[i][j] %= 9
			}
		}
	}
	return expanded
}

func Solve(lines []string) int {
	coords := make([][]int, len(lines))
	points := make(map[Point][]Point)

	for i, line := range lines {
		coords[i] = make([]int, len(line))
		for j, r := range line {
			num, _ := strconv.Atoi(string(r))
			coords[i][j] = num
		}
	}

	coords = Expand(coords)

	for i, row := range coords {
		for j, _ := range row {
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
		if point.X < len(coords[0])-1 {
			edges = append(edges, Point{X: point.X + 1, Y: point.Y})
		}
		if point.Y < len(coords)-1 {
			edges = append(edges, Point{X: point.X, Y: point.Y + 1})
		}

		points[point] = edges
	}

	costFn := func(point Point) int {
		return coords[point.Y][point.X]
	}
	path := FindPaths(coords, Point{len(coords[0]) - 1, len(coords) - 1}, Point{0, 0}, points, costFn)

	//Print(coords, len(coords), len(coords[0]), path, make(map[Point]struct{}))
	sum := 0
	for _, p := range path[1:] {
		sum += coords[p.Y][p.X]
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
