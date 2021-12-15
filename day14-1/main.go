package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Rule struct {
	Left, Right byte
	Insert      byte
}

type Insert struct {
	Pos  int
	Char byte
}

func Step(positions []byte, rules []Rule) []byte {
	inserts := make([]Insert, 0)
	for i, b := range positions[:len(positions)-1] {
		for _, rule := range rules {

			if rule.Left == b && positions[i+1] == rule.Right {
				//match!!
				inserts = append(inserts, Insert{Pos: i + 1, Char: rule.Insert})
				break
			}
		}
	}

	for i, insert := range inserts {
		pos := insert.Pos + i
		positions = append(positions[:pos+1], positions[pos:]...)
		positions[pos] = insert.Char
	}
	return positions
}

func Solve(lines []string) int {
	positions := []byte(lines[0])
	rules := make([]Rule, 0)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules = append(rules, Rule{Left: parts[0][0], Right: parts[0][1], Insert: parts[1][0]})
	}

	for i := 0; i < 10; i++ {
		positions = Step(positions, rules)
	}

	freqs := make(map[string]int)
	for _, b := range positions {
		freqs[string(b)]++
	}

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
	start := time.Now()
	defer func() {
		fmt.Println("Elapsed Time: ", time.Now().Sub(start))
	}()
	fmt.Println(Solve(ReadInput(os.Stdin)))
}
