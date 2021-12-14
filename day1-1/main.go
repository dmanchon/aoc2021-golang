package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(lines []string) int {
	increments := 0
	var previous *int
	for _, s := range lines {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
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
	solution := solve(lines)
	fmt.Println(solution)
}
