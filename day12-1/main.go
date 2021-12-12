package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindPaths(graph map[string][]string, start, end string) [][]string {
	res := make([][]string, 0)
	FindPathsRecur(graph, start, end, []string{}, &res)
	return res
}

func ShouldStopTraversing(node string, visited []string) bool {
	if strings.ToLower(node) == node {
		for _, v := range visited {
			if v == node {
				return true
			}
		}
	}
	return false
}

func FindPathsRecur(graph map[string][]string, start, end string, traversed []string, result *[][]string) {
	// Finished?
	if start == end {
		traversed = append(traversed, start)
		*result = append(*result, traversed)
		return
	} else if ShouldStopTraversing(start, traversed) {
		return
	}

	// Add current node to traversed list
	traversed = append(traversed, start)

	for _, node := range graph[start] {
		FindPathsRecur(graph, node, end, traversed, result)
	}
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
