package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseCmd(line string) (string, int) {
	re := regexp.MustCompile(`(?P<cmd>forward|down|up) (?P<amount>\d+)`)
	names := re.SubexpNames()
	matches := re.FindAllStringSubmatch(line, -1)
	m := map[string]string{}
	for i, n := range matches[0] {
		m[names[i]] = n
	}
	cmd := m["cmd"]
	amount, _ := strconv.Atoi(m["amount"])
	return cmd, amount
}

func solve(lines []string) int {
	var x, y int
	for _, line := range lines {
		cmd, amount := parseCmd(line)

		switch cmd {
		case "forward":
			x += amount
		case "up":
			y -= amount
		case "down":
			y += amount
		}
	}
	return x * y
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
