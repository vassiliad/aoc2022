package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day14/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartB(input)

	const correct_answer = 93

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
