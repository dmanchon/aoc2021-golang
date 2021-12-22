package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Position int
	Score    int
}

type Board struct {
	Turn    int
	Player1 Player
	Player2 Player
}

func Toggle(turn int) int {
	if turn == 0 {
		return 1
	}
	return 0
}

func Play(board Board, visited map[Board][2]uint64) [2]uint64 {
	var hits [2]uint64
	if hits, ok := visited[board]; ok {
		return hits
	}

	if board.Player1.Score >= 21 {
		hits[0]++
		visited[board] = hits
		return hits
	} else if board.Player2.Score >= 21 {
		hits[1]++
		visited[board] = hits
		return hits
	}

	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				value := i + j + k

				b := Board{}
				p1 := Player{}
				p2 := Player{}
				if board.Turn == 0 {
					p1.Position = (board.Player1.Position + value) % 10
					if p1.Position == 0 {
						p1.Score = board.Player1.Score + 10
					} else {
						p1.Score = board.Player1.Score + p1.Position
					}
					p2.Position = board.Player2.Position
					p2.Score = board.Player2.Score
					b.Turn = 1
					b.Player1 = p1
					b.Player2 = p2
				} else {
					p2.Position = (board.Player2.Position + value) % 10
					if p2.Position == 0 {
						p2.Score = board.Player2.Score + 10
					} else {
						p2.Score = board.Player2.Score + p2.Position
					}
					p1.Position = board.Player1.Position
					p1.Score = board.Player1.Score
					b.Turn = 0
					b.Player1 = p1
					b.Player2 = p2
				}

				res := Play(b, visited)
				hits[0] += res[0]
				hits[1] += res[1]
			}
		}
	}

	visited[board] = hits
	return hits
}

func Solve(lines []string) uint64 {
	pos1_str := strings.Split(lines[0], ": ")[1]
	pos2_str := strings.Split(lines[1], ": ")[1]

	pos1, _ := strconv.Atoi(pos1_str)
	pos2, _ := strconv.Atoi(pos2_str)

	visited := make(map[Board][2]uint64)
	board := Board{
		Turn:    0,
		Player1: Player{pos1, 0},
		Player2: Player{pos2, 0},
	}

	res := Play(board, visited)
	if res[0] > res[1] {
		return res[0]
	}

	return res[1]
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
