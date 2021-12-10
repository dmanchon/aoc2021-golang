package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Solve(lines []string) int {
	nums := make([]int, 0)
	points := map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}

	for _, line := range lines {
		open := make([]rune, 0)
	outer:
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				open = append([]rune{r}, open...)
			default:
				if open[0] != pairs[r] {
					break outer
				}
				open = open[1:]
			}
		}

		n := 0
		for _, r := range open {
			n = n*5 + points[r]
		}
		nums = append(nums, n)
	}

	sort.Ints(nums)
	return nums[len(nums)/2]
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
