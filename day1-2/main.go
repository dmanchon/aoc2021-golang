package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func window(lines []string) []int {
	result := make([]int, 0)
	var previous [3]*int

loop:
	for _, s := range lines {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		for i, prev := range previous {
			if prev == nil {
				previous[i] = &x
				continue loop
			}
		}
		sum := *previous[0] + *previous[1] + *previous[2]
		result = append(result, sum)

		previous[0] = previous[1]
		previous[1] = previous[2]
		previous[2] = &x
	}

	sum := *previous[0] + *previous[1] + *previous[2]
	result = append(result, sum)
	return result
}

func solve(lines []int) int {
	increments := 0
	var previous *int
	for _, x := range lines {
		x := x
		if previous != nil && (x-*previous) > 0 {
			increments += 1
		}
		previous = &x
	}
	return increments
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sums := window(lines)
	solution := solve(sums)
	fmt.Println(solution)
}
