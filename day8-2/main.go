package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type RuneSet map[rune]struct{}

func BuildSet(s string) RuneSet {
	set := make(RuneSet)
	for _, r := range s {
		set[r] = struct{}{}
	}
	return set
}

func (s RuneSet) Equal(s1 RuneSet) bool {
	return len(s.Difference(s1)) == 0 && (len(s) == len(s1))
}

func (s RuneSet) String() string {
	str := make([]rune, 0)
	for k, _ := range s {
		str = append(str, k)
	}
	return string(str)
}

func (s RuneSet) Union(s1 RuneSet) RuneSet {
	union := make(RuneSet)
	for k, _ := range s {
		if v, ok := s1[k]; ok {
			union[k] = v
		}
	}
	return union
}

func (s RuneSet) Difference(s1 RuneSet) RuneSet {
	diff := make(RuneSet)
	for k, _ := range s {
		if v, ok := s1[k]; !ok {
			diff[k] = v
		}
	}
	return diff
}

func (s RuneSet) Len() int {
	return len(s)
}

func Filter(list []RuneSet, size int) []RuneSet {
	res := make([]RuneSet, 0)
	for _, s := range list {
		if s.Len() == size {
			res = append(res, s)
		}
	}
	return res
}

func Solve(lines []string) int {
	sums := 0

	results := make(chan int, len(lines))
	var wg sync.WaitGroup

	wg.Add(len(lines))
	for _, line := range lines {
		go func(line string) {
			results <- SolveLine(line)
			wg.Done()
		}(line)
	}
	wg.Wait()
	close(results)

	for r := range results {
		sums += r
	}
	return sums
}

func SolveLine(line string) int {
	split := strings.Split(line, "|")
	digits := strings.Split(split[0], " ")

	var zero, one, two, three, four, five, six, seven, eight, nine RuneSet
	list := make([]RuneSet, len(digits))
	for _, digit := range digits {
		s := BuildSet(digit)
		list = append(list, s)
	}

	one = Filter(list, 2)[0]
	eight = Filter(list, 7)[0]
	seven = Filter(list, 3)[0]
	four = Filter(list, 4)[0]

	for _, candidate := range Filter(list, 6) {
		size := candidate.Difference(one).Len()
		if size == 5 {
			six = candidate
		}
	}

	for _, candidate := range Filter(list, 5) {
		size := candidate.Difference(one).Len()
		if size == 3 {
			three = candidate
			break
		}
	}

	for _, candidate := range Filter(list, 6) {
		if candidate.Equal(six) {
			continue
		}
		size := three.Difference(candidate).Len()
		if size == 0 {
			nine = candidate
		} else {
			zero = candidate
		}
	}

	for _, candidate := range Filter(list, 5) {
		if candidate.Equal(three) {
			continue
		}
		size := six.Difference(candidate).Len()
		if size == 2 {
			two = candidate
		} else if size == 1 {
			five = candidate
		}
	}

	number := make([]rune, 0)
	for _, digit := range strings.Split(split[1], " ")[1:] {
		s := BuildSet(digit)

		if zero.Equal(s) {
			number = append(number, '0')
		} else if one.Equal(s) {
			number = append(number, '1')
		} else if two.Equal(s) {
			number = append(number, '2')
		} else if three.Equal(s) {
			number = append(number, '3')
		} else if four.Equal(s) {
			number = append(number, '4')
		} else if five.Equal(s) {
			number = append(number, '5')
		} else if six.Equal(s) {
			number = append(number, '6')
		} else if seven.Equal(s) {
			number = append(number, '7')
		} else if eight.Equal(s) {
			number = append(number, '8')
		} else if nine.Equal(s) {
			number = append(number, '9')
		} else {
			number = append(number, '?')
		}

	}

	n, _ := strconv.Atoi(string(number))

	return n
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
