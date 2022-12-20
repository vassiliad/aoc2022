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
	metadata, solution := PartB(input)

	// VV: As long as 0 is in the correct index wrt to other numbers, we're fine
	expected := []int{0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153}
	produced := metadata.ToSlice()

	fmt.Println(produced)

	if !reflect.DeepEqual(expected, produced) {
		t.Fatal("Expected mixed numbers to be", expected, "but they were", produced)
	}

	const correct_answer = 1623178306
	correct_grove := [3]int{811589153, 2434767459, -1623178306}

	grove := metadata.ToGroveCoords()

	if grove != correct_grove {
		t.Fatal("Expected Grove to be", correct_grove, "but it was", grove)
	}

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}

}

func TestStep0(t *testing.T) {
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

	const key = 811589153

	metadata := Metadata{}

	metadata.Prepare(input, key)

	fmt.Println("BEFORE", metadata.ToSlice())
	metadata.MoveNumber(0)

	expected := []int{0, 811589153, 3246356612, 1623178306, -2434767459, 2434767459, -1623178306}
	produced := metadata.ToSlice()

	fmt.Println(metadata.Order[0].Value, "Produced", produced)

	if !reflect.DeepEqual(expected, produced) {
		t.Fatal("Expected mixed numbers 0 to be", expected, "but they were", produced)
	}

	metadata.MoveNumber(1)

	expected = []int{0, 1623178306, 811589153, 3246356612, -2434767459, 2434767459, -1623178306}
	produced = metadata.ToSlice()

	fmt.Println(metadata.Order[1].Value, "Produced", produced)

	if !reflect.DeepEqual(expected, produced) {
		t.Fatal("Expected mixed numbers 1 to be", expected, "but they were", produced)
	}

	metadata.MoveNumber(2)

	expected = []int{0, -2434767459, 1623178306, 811589153, 3246356612, 2434767459, -1623178306}
	produced = metadata.ToSlice()

	fmt.Println(metadata.Order[2].Value, "Produced", produced)

	if !reflect.DeepEqual(expected, produced) {
		t.Fatal("Expected mixed numbers 1 to be", expected, "but they were", produced)
	}
}
