package main

import (
	"fmt"
	"testing"

	"github.com/vassiliad/aoc2022/day09/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartA(input)

	const correct_answer = 13

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestDiagonnalMove(t *testing.T) {
	head_x, head_y, tail_x, tail_y := SimulateRope(
		&utilities.Move{Direction: utilities.DirectionUp, Steps: 1}, nil,
		4, 1, 3, 0,
	)

	fmt.Printf("Head (%d, %d), Tail (%d, %d)", head_x, head_y, tail_x, tail_y)

	if head_x != 4 {
		t.Fatal("head_x", head_x, 4)
	}

	if head_y != 2 {
		t.Fatal("head_y", head_y, 2)
	}

	if tail_x != 4 {
		t.Fatal("tail_x", tail_x, 4)
	}

	if tail_y != 1 {
		t.Fatal("tail_y", tail_y, 1)
	}
}
