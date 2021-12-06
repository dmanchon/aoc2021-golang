package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(lines []string) int {

	population := make([]int, 0)
	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		population = append(population, n)
	}

	for day := 0; day < 81; day++ {
		newborns := 0
		for i, fish := range population {
			if fish == 0 {
				population[i] = 6
				newborns++
			} else {
				population[i]--
			}
		}
		for i := 0; i < newborns; i++ {
			population = append(population, 8)
		}
	}

	return len(population)
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
