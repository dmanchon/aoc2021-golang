package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(lines []string) int {

	num := 0
	for _, line := range lines {
		digits := strings.Split(strings.Split(line, "|")[1], " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				num++
			case 3:
				num++
			case 7:
				num++
			case 4:
				num++
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
