package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Coefficient(x, y int, image [][]rune, fill rune) int64 {
	digits := make([]rune, 0)
	for _, i := range []int{-1, 0, +1} {
		for _, j := range []int{-1, 0, +1} {
			var value rune
			iy := i + y
			ix := j + x
			if iy < 0 || iy > len(image)-1 || ix < 0 || ix > len(image[iy])-1 {
				value = fill
			} else {
				value = image[iy][ix]
			}
			digits = append(digits, value)
		}
	}
	n, _ := strconv.ParseInt(string(digits), 2, 64)
	return n
}

func Step(in [][]rune, coeff []rune, fill rune) [][]rune {

	extend := 3
	new := make([][]rune, len(in)+(extend*2))
	new = append(new[:extend], in...)
	new = append(new[:extend+len(in)], new[:extend]...)

	for i := range new {
		new[i] = make([]rune, len(in)+(extend*2))
		if i < extend || i > len(new)-(extend+1) {
			for k := range new[i] {
				if new[i][k] == rune(0) {
					new[i][k] = fill
				}
			}
			continue
		}
		new[i] = append(new[i][:extend], in[i-extend]...)
		new[i] = append(new[i][:extend+len(in[i-extend])], new[i][:extend]...)

		for k := range new[i] {
			if new[i][k] == rune(0) {
				new[i][k] = fill
			}
		}
	}

	out := make([][]rune, len(new))
	for i := range new {
		out[i] = make([]rune, len(new[i]))
		for j := range new[i] {
			x := Coefficient(j, i, new, fill)
			r := rune(coeff[int(x)])
			if r == '.' {
				out[i][j] = '0'
			} else {
				out[i][j] = '1'
			}
		}
	}
	return out
}

func Print(in [][]rune) {
	for i := range in {
		s := string(in[i])
		s = strings.Replace(s, "0", ".", -1)
		s = strings.Replace(s, "1", "#", -1)
		fmt.Println(s)
	}
}
func Solve(lines []string) int {
	enhancement := lines[0]

	image := make([][]rune, len(lines[2:]))
	for i, line := range lines[2:] {
		image[i] = make([]rune, len(line))
		for j, r := range line {
			if r == '.' {
				image[i][j] = '0'
			} else {
				image[i][j] = '1'
			}
		}
	}

	for i := 0; i < 50; i++ {
		var fill rune

		// flip bits
		if i%2 == 0 {
			fill = '0'
		} else {
			if enhancement[0] == '#' {
				fill = '1'
			}
		}
		image = Step(image, []rune(enhancement), fill)
	}

	sum := 0
	for _, l := range image {
		for _, r := range l {
			if r == '1' {
				sum++
			}
		}
	}

	return sum
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
