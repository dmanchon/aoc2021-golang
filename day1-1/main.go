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
		if previous == nil {
			fmt.Printf("%d (N/A - no previous measurement)\n", x)
		} else if (x - *previous) > 0 {
			fmt.Printf("%d (increased)\n", x)
			increments += 1
		} else {
			fmt.Printf("%d (decreased)\n", x)
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
