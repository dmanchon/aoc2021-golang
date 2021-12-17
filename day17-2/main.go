package main

import (
	"fmt"
	"time"
)

func CheckTrajectory(vx, vy, x0, x1, y0, y1 int) (bool, int) {
	maxY := 0
	x, y := 0, 0

	for {
		x += vx
		y += vy

		if y > maxY {
			maxY = y
		}

		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		vy--

		// on target
		if x >= x0 && x <= x1 && y >= y0 && y <= y1 {
			return true, maxY
		}

		// will never reach target
		if x > x1 || y < y0 {
			return false, -1
		}
	}
}

func EtTuBrute(x0, x1, y0, y1 int) int {
	num := make(map[Point]struct{})
	for i := -300; i < 300; i++ {
		for j := -10; j < 300; j++ {
			match, _ := CheckTrajectory(j, i, x0, x1, y0, y1)
			if match {
				num[Point{i, j}] = struct{}{}
			}
		}
	}
	return len(num)
}

type Point struct {
	x, y int
}

func Solve() int {
	// also harcoding the input.
	return EtTuBrute(230, 283, -107, -57)
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("- Elapsed Time: ", time.Since(start))
	}()
	fmt.Println("- Solution: ", Solve())
}
