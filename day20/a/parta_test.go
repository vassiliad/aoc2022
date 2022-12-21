package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/vassiliad/aoc2022/day20/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	1
	2
	-3
	3
	-2
	0
	4`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	metadata, solution := PartA(input)

	const correct_answer = 3
	correct_grove := [3]int{4, -3, 2}

	grove := metadata.ToGroveCoords()

	if grove != correct_grove {
		t.Fatal("Expected Grove to be", correct_grove, "but it was", grove)
	}

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}

	// VV: As long as 0 is in the correct index wrt to other numbers, we're fine
	expected := []int{0, 3, -2, 1, 2, -3, 4}
	produced := metadata.ToSlice()

	fmt.Println(produced)

	if !reflect.DeepEqual(expected, produced) {
		t.Fatal("Expected mixed numbers to be", expected, "but they were", produced)
	}

}
