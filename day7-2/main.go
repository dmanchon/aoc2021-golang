package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Fuel(n int) int {
	n = int(math.Abs(float64(n))) + 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func Solve(lines []string) int {
	pos := make([]int, 0)
	min, max := 0, 0
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		pos = append(pos, n)

		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	distance := make([]int, 0)
	for n := min; n < max; n++ {
		sum := 0
		for _, x := range pos {
			fuel := Fuel(x - n)
			sum += fuel
		}
		distance = append(distance, sum)
	}

	solution := distance[0]
	for _, x := range distance {
		if x < solution {
			solution = x
		}
	}

	return solution
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
