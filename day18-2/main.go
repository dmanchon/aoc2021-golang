package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Number struct {
	Value       *int
	Left, Right *Number
	Depth       int
}

type Element struct {
	Value int
	Depth int
}

func Flatten(n *Number) []Element {
	lst := make([]Element, 0)

	if n.Left != nil {
		lst = append(lst, Flatten(n.Left)...)
	}
	if n.Right != nil {
		lst = append(lst, Flatten(n.Right)...)
	}

	if n.Value != nil {
		lst = append(lst, Element{Value: *n.Value, Depth: n.Depth})
	}
	return lst
}

func AllZeroDepth(number []Element) bool {
	for _, n := range number {
		if n.Depth > 0 {
			return false
		}
	}
	return true
}

func Magnitude(num []Element) int {

	number := make([]Element, len(num))
	copy(number, num)
	i := 0
outer:
	for {
		for !AllZeroDepth(number) {
			if number[i].Depth > 0 && number[i+1].Depth == number[i].Depth {
				val := 3*number[i].Value + 2*number[i+1].Value
				depth := number[i].Depth - 1
				number = append(number[:i], number[i+1:]...)
				number[i] = Element{Value: val, Depth: depth}
				i = 0
				continue outer
			}

			i++
			if i > len(number)-1 {
				i = 0
			}

		}

		return number[0].Value*3 + number[1].Value*2
	}
}

func NewNumber(num_str string) *Number {
	var recur func(s string, depth int) *Number
	recur = func(s string, depth int) *Number {
		var opened int

		// Value
		if s[0] != '[' {
			n, _ := strconv.Atoi(s)
			return &Number{Value: &n, Depth: depth}
		}

		// Pair
		inner := s[1 : len(s)-1]
		for i, c := range inner {

			if c == '[' {
				opened++
			} else if c == ']' {
				opened--
			}
			//
			if opened == 0 {
				return &Number{
					Depth: depth,
					Left:  recur(inner[:i+1], depth+1),
					Right: recur(inner[i+2:], depth+1),
				}
			}
		}
		return nil
	}
	return recur(num_str, -1)
}

func Sum(f1, f2 []Element) []Element {
	// each flatten list gets one level deeper
	flatten := make([]Element, 0)
	flatten = append(f1, f2...)
	for i := range flatten {
		flatten[i].Depth++
	}
outer:
	for {
		i := 0
		for i < len(flatten) {
			p1 := flatten[i]
			// Explodes?
			if p1.Depth > 3 {
				p2 := flatten[i+1]
				flatten = append(flatten[:i], flatten[i+1:]...)
				flatten[i] = Element{Value: 0, Depth: 3}

				if i > 0 {
					flatten[i-1].Value += p1.Value
				}

				if i < len(flatten)-1 {
					flatten[i+1].Value += p2.Value
				}

				i = 0
				continue outer
			}
			i++
		}
		i = 0
		for i < len(flatten) {
			p1 := flatten[i]
			// Split
			if p1.Value > 9 {
				flatten[i] = Element{Value: p1.Value/2 + p1.Value%2, Depth: p1.Depth + 1}
				flatten = append(flatten[:i+1], flatten[i:]...)
				flatten[i] = Element{Value: p1.Value / 2, Depth: p1.Depth + 1}
				i = 0
				continue outer
			}
			i++
		}
		break
	}
	return flatten
}

func Solve(lines []string) int {

	numbers := make([][]Element, 0)
	for _, line := range lines {
		numbers = append(numbers, Flatten(NewNumber(line)))
	}

	max := 0
	for i := range numbers {
		for j := range numbers {
			if i != j {
				// i dont understand yet why i need to do this copy!!! :(
				n1 := make([]Element, len(numbers[i]))
				n2 := make([]Element, len(numbers[j]))
				copy(n1, numbers[i])
				copy(n2, numbers[j])

				magnitude := Magnitude(Sum(n1, n2))
				if magnitude > max {
					max = magnitude
				}
			}
		}
	}
	return max
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
