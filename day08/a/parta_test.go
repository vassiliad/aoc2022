package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day08/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	30373
	25512
	65332
	33549
	35390
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartA(input)

	const correct_answer = 21

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
