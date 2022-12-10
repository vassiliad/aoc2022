package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day04/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	const num_items = 6

	if len(input) != num_items {
		t.Fatal("Expected", num_items, "items but read", len(input))
	}

	solution := PartA(input)

	const correct_answer = 2

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
