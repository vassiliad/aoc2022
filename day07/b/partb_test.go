package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day07/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartB(&input)

	const correct_answer = uint64(24933642)

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
