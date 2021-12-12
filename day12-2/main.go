package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindPaths(graph map[string][]string, start, end string) [][]string {
	res := make([][]string, 0)
	var recur func(start string, traversed []string)

	recur = func(start string, traversed []string) {
		traversed = append(traversed, start)
		if start == end {
			res = append(res, traversed)
			return
		} else if ShouldStopTraversing(start, traversed[:len(traversed)-1]) {
			return
		}
		for _, node := range graph[start] {
			recur(node, traversed)
		}
	}
	recur(start, []string{})
	return res
}

func ShouldStopTraversing(node string, traversed []string) bool {
	freqs := make(map[string]int)

	freqs[node] = 1
	repeated := 0

	if strings.ToLower(node) != node {
		return false
	}

	for _, v := range traversed {
		if strings.ToLower(v) != v {
			continue
		}

		if v == "start" && v == node {
			return true
		}

		freqs[v]++
		if freqs[v] > 1 {
			repeated++
		}
		if repeated > 1 {
			return true
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
