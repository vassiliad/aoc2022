package main

import (
	"testing"
)

func TestSmall4(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	solution := PartA(input)

	const correct_answer = 7

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall0(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	solution := PartA(input)

	const correct_answer = 5

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall1(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"

	solution := PartA(input)

	const correct_answer = 6

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall2(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	solution := PartA(input)

	const correct_answer = 10

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestSmall3(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	solution := PartA(input)

	const correct_answer = 11

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
