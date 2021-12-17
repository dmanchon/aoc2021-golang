package main

import (
	"fmt"
	"math"
	"time"
)

func CheckTrajectory(vx, vy, x0, x1, y0, y1 int) (bool, int) {
	maxY := math.MinInt
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
		if x > x1 || y < y1 {
			return false, -1
		}
	}
}

func EtTuBrute(x0, x1, y0, y1 int) int {
	max := math.MinInt
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			match, m := CheckTrajectory(i, j, x0, x1, y0, y1)
			if match && m > max {
				max = m
			}
		}
	}
	return max
}

func Solve() int {
	// also harcoding the input.
	return EtTuBrute(230, 238, -107, -57)
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("- Elapsed Time: ", time.Since(start))
	}()
	fmt.Println("- Solution: ", Solve())
}
