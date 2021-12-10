package main

import (
	"bufio"
	"fmt"
	"os"
)

func Solve(lines []string) int {
	num := 0
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

outer:
	for _, line := range lines {
		open := make([]rune, 0)
		for _, r := range line {

			switch r {
			case '(', '[', '{', '<':
				open = append(open, r)
			case ')':
				if open[len(open)-1] != '(' {
					num += points[r]
					continue outer
				}
				open = open[:len(open)-1]
			case ']':
				if open[len(open)-1] != '[' {
					num += points[r]
					continue outer
				}
				open = open[:len(open)-1]

			case '}':
				if open[len(open)-1] != '{' {
					num += points[r]
					continue outer
				}
				open = open[:len(open)-1]

			case '>':
				if open[len(open)-1] != '<' {
					num += points[r]
					continue outer
				}
				open = open[:len(open)-1]

			}
		}

	}

	return num
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
