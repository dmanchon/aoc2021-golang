package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Range struct {
	Start, End int
}

type Action struct {
	On bool
	X  Range
	Y  Range
	Z  Range
}

type Coord struct {
	X, Y, Z int
}

func OutOfRange(r Range, Min, Max int) bool {
	if r.Start < Min || r.End > Max {
		return true
	}
	return false
}

func Solve(lines []string) int {

	actions := make([]Action, 0)
	for _, line := range lines {
		re := regexp.MustCompile(`(?P<action>on|off) x=(?P<x0>-?\d+)..(?P<x1>-?\d+),y=(?P<y0>-?\d+)..(?P<y1>-?\d+),z=(?P<z0>-?\d+)..(?P<z1>-?\d+)`)
		names := re.SubexpNames()
		matches := re.FindAllStringSubmatch(line, -1)
		m := map[string]string{}
		for i, n := range matches[0] {
			m[names[i]] = n
		}

		on := m["action"] == "on"
		x0, _ := strconv.Atoi(m["x0"])
		x1, _ := strconv.Atoi(m["x1"])
		y0, _ := strconv.Atoi(m["y0"])
		y1, _ := strconv.Atoi(m["y1"])
		z0, _ := strconv.Atoi(m["z0"])
		z1, _ := strconv.Atoi(m["z1"])

		actions = append(actions, Action{on, Range{x0, x1}, Range{y0, y1}, Range{z0, z1}})
	}

	state := make(map[Coord]struct{})

	for _, action := range actions {
		if OutOfRange(action.X, -50, 50) || OutOfRange(action.Y, -50, 50) || OutOfRange(action.Z, -50, 50) {
			continue
		}
		for x := action.X.Start; x <= action.X.End; x++ {
			for y := action.Y.Start; y <= action.Y.End; y++ {
				for z := action.Z.Start; z <= action.Z.End; z++ {
					coord := Coord{x, y, z}

					if action.On {
						state[coord] = struct{}{}
					} else {
						if _, ok := state[coord]; ok {
							delete(state, coord)
						}
					}
				}
			}
		}
	}

	return len(state)
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
