package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day03/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	vJrwpWtwJgWrhcsFMMfFFhFp
    jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
    PmmdzqPrVvPwwTWBwg
    wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
    ttgJtRGJQctTZtZT
    CrZsJsPPZsGzwwsLwLmpwMDw`

	compartments, _ := utilities.ReadString(small)

	const num_items = 6

	if len(compartments) != num_items {
		t.Fatal("Expected", num_items, "items but read", len(compartments))
	}

	solution := PartA(compartments)

	const correct_answer = 157

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
