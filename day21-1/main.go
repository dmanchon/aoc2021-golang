package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Dice struct {
	Value int
	Rolls int
}

func (d *Dice) Roll() (result int) {
	d.Rolls++
	result = d.Value

	d.Value++
	if d.Value > 100 {
		d.Value = 1
	}

	return result
}

func Solve(lines []string) int {
	pos1_str := strings.Split(lines[0], ": ")[1]
	pos2_str := strings.Split(lines[1], ": ")[1]

	pos1, _ := strconv.Atoi(pos1_str)
	pos2, _ := strconv.Atoi(pos2_str)

	score1, score2 := 0, 0
	dice := Dice{Value: 1}
	loser := 0

	for {
		//player1 rolls
		roll := dice.Roll() + dice.Roll() + dice.Roll()
		pos1 = (pos1 + roll) % 10
		if pos1 == 0 {
			score1 += 10
		} else {
			score1 += pos1
		}
		if score1 >= 1000 {
			loser = score2
			break
		}

		// plauer2 rolls
		roll = dice.Roll() + dice.Roll() + dice.Roll()
		pos2 = (pos2 + roll) % 10
		if pos2 == 0 {
			score2 += 10
		} else {
			score2 += pos2
		}

		if score2 >= 1000 {
			loser = score1
			break
		}
	}
	return loser * dice.Rolls

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
		fmt.Println("- Elapsed Time: ", time.Since(start))
	}()
	fmt.Println("- Solution: ", Solve(ReadInput(os.Stdin)))
}
