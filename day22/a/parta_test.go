package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day22/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `        
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	instr_parsed := input.InstructionsToString()
	instr_expected := "10R5L5R10L4R5L5"

	if instr_parsed != "10R5L5R10L4R5L5" {
		t.Fatalf("Expected to parse instructions %s but parsed %s", instr_expected, instr_parsed)
	}

	solution := PartA(input)

	const correct_answer = 6032

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
