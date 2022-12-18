package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day18/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	2,2,2
	1,2,2
	3,2,2
	2,1,2
	2,3,2
	2,2,1
	2,2,3
	2,2,4
	2,2,6
	1,2,5
	3,2,5
	2,1,5
	2,3,5`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartA(input)

	const correct_answer = 64

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
