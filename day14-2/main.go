package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Rule struct {
	Left, Right byte
	Insert      byte
}

type Pair struct {
	Left, Right byte
}

func Step(pairs map[string]int, rules []Rule) map[string]int {
	update := make(map[string]int)
	for k, v := range pairs {
		update[k] = v
	}

	for p, count := range pairs {
		for _, rule := range rules {
			if rule.Left == p[0] && p[1] == rule.Right {
				//match!!
				update[string([]byte{p[0], rule.Insert})] += count
				update[string([]byte{rule.Insert, p[1]})] += count
				update[p] -= count
				break
			}
		}
	}
	return update
}

func Solve(lines []string) int {

	pairs := make(map[string]int)
	var delim byte = 0x00
	// head
	pairs[string([]byte{delim, lines[0][0]})] = 1
	// tail
	pairs[string([]byte{lines[0][len(lines[0])-1], byte(delim)})] = 1

	for i, b := range lines[0][:len(lines[0])-1] {
		s := []byte{byte(b), lines[0][i+1]}
		pairs[string(s)]++
	}

	rules := make([]Rule, 0)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules = append(rules, Rule{Left: parts[0][0], Right: parts[0][1], Insert: parts[1][0]})
	}

	for i := 0; i < 40; i++ {
		pairs = Step(pairs, rules)

	}

	freqs := make(map[byte]int)
	for p, count := range pairs {
		// 'freqs[p[1]] += count' workds too
		freqs[p[0]] += count

	}

	// we dont care about delimiters
	delete(freqs, delim)
	max := 0
	min := math.MaxInt
	for _, count := range freqs {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max - min
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
