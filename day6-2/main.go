package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(lines []string) int {

	population := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		7: 0,
		8: 0,
	}

	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		population[n]++
	}

	for day := 0; day < 256; day++ {
		Δ := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			7: 0,
			8: 0,
		}

		for age, number := range population {
			if age == 0 {
				Δ[6] += number
				Δ[8] += number
				Δ[0] -= number
				continue
			}
			Δ[age] -= number
			Δ[age-1] += number
		}

		for age, number := range Δ {
			population[age] += number
		}

	}

	sum := 0
	for _, n := range population {
		sum += n
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
