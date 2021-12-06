package main

import (
	"os"
	"testing"
)

func BenchmarkExercise(b *testing.B) {
	file, _ := os.Open("input.txt")
	lines := ReadInput(file)
	for i := 0; i < b.N; i++ {
		Solve(lines)
	}
}

func TestExercise(t *testing.T) {
	file, _ := os.Open("input.txt")
	lines := ReadInput(file)
	solution := Solve(lines)
	if solution != 1675781200288 {
		t.Errorf("Expected '1675781200288' got '%d'", solution)
	}

}
