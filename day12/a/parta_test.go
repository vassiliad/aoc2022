package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day12/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	Sabqponm
	abcryxxl
	accszExk
	acctuvwj
	abdefghi`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	if input.Width != 8 {
		t.Fatal("Expected World.Width to be 8 but it was", input.Width)
	}

	if input.Height != 5 {
		t.Fatal("Expected World.Height to be 5 but it was", input.Height)
	}

	solution := PartA(input)

	const correct_answer = 31

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
