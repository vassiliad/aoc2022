package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day05/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	const num_stacks = 3
	const num_orders = 4

	if len(input.Stacks) != num_stacks {
		t.Fatal("Expected", num_stacks, "stacks but read", len(input.Stacks))
	}

	if len(input.Orders) != num_orders {
		t.Fatal("Expected", num_orders, "orders but read", len(input.Orders))
	}

	solution := PartB(input)

	const correct_answer = "MCD"

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
