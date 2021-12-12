package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindPaths(graph map[string][]string, start, end string) [][]string {
	res := make([][]string, 0)
	var recur func(start, end string, traversed []string)

	recur = func(start, end string, traversed []string) {
		traversed = append(traversed, start)
		if start == end {
			res = append(res, traversed)
			return
		} else if ShouldStopTraversing(start, traversed) {
			return
		}
		for _, node := range graph[start] {
			recur(node, end, traversed)
		}

	}
	recur(start, end, []string{})
	return res
}

func ShouldStopTraversing(node string, traversed []string) bool {
	if strings.ToLower(node) == node {
		for _, v := range traversed[:len(traversed)-1] {
			if v == node {
				return true
			}
		}
	}
	return false
}

func Solve(lines []string) int {
	graph := make(map[string][]string)
	for _, line := range lines {
		s := strings.Split(line, "-")
		vertices, ok := graph[s[0]]
		if !ok {
			vertices = make([]string, 0)
		}
		vertices = append(vertices, s[1])
		graph[s[0]] = vertices

		vertices, ok = graph[s[1]]
		if !ok {
			vertices = make([]string, 0)
		}
		vertices = append(vertices, s[0])
		graph[s[1]] = vertices
	}
	paths := FindPaths(graph, "start", "end")
	return len(paths)
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
