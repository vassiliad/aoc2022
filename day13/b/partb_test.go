package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day13/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	[1,1,3,1,1]
	[1,1,5,1,1]

	[[1],[2,3,4]]
	[[1],4]

	[9]
	[[8,7,6]]

	[[4,4],4,4]
	[[4,4],4,4,4]

	[7,7,7,7]
	[7,7,7]

	[]
	[3]

	[[[]]]
	[[]]

	[1,[2,[3,[4,[5,6,7]]]],8,9]
	[1,[2,[3,[4,[5,6,0]]]],8,9]`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartB(input)

	const correct_answer = 140

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
