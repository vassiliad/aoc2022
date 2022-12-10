package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day01/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
		1000
		2000
		3000
		
		4000
		
		5000
		6000

		7000
		8000
		9000
		
		10000`

	elves, _ := utilities.ReadString(small)

	if len(elves) != 5 {
		t.Fatal("Expected 5 elves but read", len(elves))
	}

	solution := PartA(elves)

	if solution != 24000 {
		t.Fatal("Expected answer to be 24000 but it was", solution)
	}
}
