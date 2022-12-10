package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day02/a/utilities"
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

	solution := PartA(rounds)

	if solution != 15 {
		t.Fatal("Expected answer to be 15 but it was", solution)
	}
}
