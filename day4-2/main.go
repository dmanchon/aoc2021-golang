package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	Id      int
	Columns [5]map[int]bool
	Rows    [5]map[int]bool
}

type Bingo struct {
	Numbers []int
	Boards  []Board
}

func (b Board) String() string {
	repr := fmt.Sprintf("Board no. %d\n", b.Id)

	for i, row := range b.Rows {
		repr = repr + fmt.Sprintf("[%d] -> %v\n", i, row)
	}
	return repr
}

func (b Board) AddNumber(number int) {
	for _, r := range b.Rows {
		if _, ok := r[number]; ok {
			r[number] = true
		}

	}
	for _, c := range b.Columns {
		if _, ok := c[number]; ok {
			c[number] = true
		}
	}
}

func (b Board) Sum() int {
	var sum int
	for _, c := range b.Columns {
		for k, v := range c {
			if !v {
				sum += k
			}
		}
	}
	return sum
}

func every(s map[int]bool) bool {
	for _, v := range s {
		if !v {
			return false
		}
	}
	return true
}

func (b Board) Winner() bool {
	for _, r := range b.Rows {
		if every(r) {
			return true
		}
	}

	for _, c := range b.Columns {
		if every(c) {
			return true
		}
	}
	return false
}

func (b Bingo) Play() int64 {
	winners := 0
	for _, number := range b.Numbers {
		for _, board := range b.Boards {
			// ignore the past winners
			if board.Winner() {
				continue
			}
			board.AddNumber(number)

			if board.Winner() {
				winners++
			}

			if len(b.Boards) == winners {
				return int64(board.Sum() * number)
			}
		}
	}
	return 0
}

func NewBoard(id int, lines []string) *Board {
	var columns [5]map[int]bool
	var rows [5]map[int]bool

	for i := 0; i < 5; i++ {
		columns[i] = make(map[int]bool)
		rows[i] = make(map[int]bool)
	}

	for i, row := range lines {
		re := regexp.MustCompile(`(\d+)[\s+]{0,}`)

		for j, digit := range re.FindAllStringSubmatch(row, -1) {
			n, _ := strconv.Atoi(digit[1])
			columns[j][n] = false
			rows[i][n] = false
		}

	}
	board := &Board{
		Id:      id,
		Columns: columns,
		Rows:    rows,
	}
	return board
}

func NewBingo(lines []string) *Bingo {
	var numbers []int
	var boards []Board

	for _, s := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	p := 0
	for {
		a := p*5 + 2 + p
		b := a + 5

		if b > len(lines) {
			break
		}

		board := NewBoard(p, lines[a:b])
		boards = append(boards, *board)
		p++
	}

	bingo := &Bingo{
		Numbers: numbers,
		Boards:  boards,
	}

	return bingo
}

func solve(lines []string) int {
	bingo := NewBingo(lines)
	return int(bingo.Play())
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
