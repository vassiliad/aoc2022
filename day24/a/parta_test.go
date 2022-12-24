package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day24/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `        
	#.######
	#>>.<^<#
	#.<..<<#
	#>v.><>#
	#<^v^^>#
	######.#	
	`
	/*
		#.######
		#>>.<^<#
		#.<..<<#
		#>v.><>#
		#<^v^^>#
		######.#	*/
	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartA(input)

	const correct_answer = 18

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
