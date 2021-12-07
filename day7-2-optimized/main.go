package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Fuel(n int) int {
	n = int(math.Abs(float64(n)))
	return (n * (n + 1)) / 2
}

type Crab struct {
	Pos     int
	SumDist int
}

func Solve(lines []string) int {
	pos := make([]int, 0)
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		pos = append(pos, n)
	}

	distances := make([]Crab, len(pos))
	for i, n := range pos {
		distances[i].Pos = i
		for _, m := range pos {
			distances[i].SumDist = distances[i].SumDist + Fuel(m-n)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].SumDist < distances[j].SumDist
	})

	a, b := distances[1].Pos, distances[2].Pos
	if a > b {
		a, b = b, a
	}

	min := math.MaxInt64
	for i := a; i < b; i++ {
		sum := 0
		for _, m := range pos {
			sum += Fuel(i - m)
		}
		if sum < min {
			min = sum
		}
	}

	return min
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
