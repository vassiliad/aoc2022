package main

import (
	"testing"
)

func TestSmall4(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	solution := PartB(input)

	const correct_answer = 19

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall0(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	solution := PartB(input)

	const correct_answer = 23

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall1(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	solution := PartB(input)

	const correct_answer = 23

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall2(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	solution := PartB(input)

	const correct_answer = 29

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall3(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	solution := PartB(input)

	const correct_answer = 26

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
