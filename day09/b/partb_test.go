package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day09/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	R 5
	U 8
	L 8
	D 3
	R 17
	D 10
	L 25
	U 20
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartB(input)

	const correct_answer = 36

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
