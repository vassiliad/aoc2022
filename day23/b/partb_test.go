package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day23/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	....#..
	..###.#
	#...#.#
	.#...##
	#.###..
	##.#.##
	.#..#..
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartB(input)

	const correct_answer = 20

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestRounds(t *testing.T) {
	small := `        
	....#..
	..###.#
	#...#.#
	.#...##
	#.###..
	##.#.##
	.#..#..	
	`

	group, _ := utilities.ReadString(small)

	w := Work{
		Group:        *group,
		OrderedElves: []utilities.Point{},
		Neighbours:   []map[int]int8{},
	}
	w.Init()

	w.ActRound()

	report := w.Draw(true)

	expected := `Round 1, (-1, 1) -> (7, -7)
.....#...
...#...#.
.#..#.#..
.....#..#
..#.#.##.
#..#.#...
#.#.#.##.
.........
..#..#...
`

	if report != expected {
		t.Fatal("Expected answer to be\n", expected, "\nbut it was\n", report)
	}

	w.ActRound()

	report = w.Draw(false)

	expected = `Round 2, (-2, 1) -> (8, -7)
......#....
...#.....#.
..#..#.#...
......#...#
..#..#.#...
#...#.#.#..
...........
.#.#.#.##..
...#..#....
`

	if report != expected {
		t.Fatal("Expected answer to be\n", expected, "\nbut it was\n", report)
	}

}
