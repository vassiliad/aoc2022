package main

import (
	"testing"

	utilities "github.com/vassiliad/aoc2022/day02/b/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	A Y
	B X
	C Z`

	rounds, _ := utilities.ReadString(small)

	if len(rounds) != 3 {
		t.Fatal("Expected 3 rounds but read", len(rounds))
	}

	solution := PartB(rounds)

	const correct = 12
	if solution != correct {
		t.Fatal("Expected answer to be", correct, "but it was", solution)
	}
}
