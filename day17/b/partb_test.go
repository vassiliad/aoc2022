package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day17/a/utilities"
)

func TestSmall(t *testing.T) {
	small := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartB(input)

	const correct_answer = 1514285714288

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution, "delta is", correct_answer-solution)
	}
}
