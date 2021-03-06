package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func empty() map[int]int {
	return map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		7: 0,
		8: 0,
	}
}
func Solve(lines []string) int {

	population := empty()

	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		population[n]++
	}

	for day := 0; day < 256; day++ {
		Δ := empty()
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
