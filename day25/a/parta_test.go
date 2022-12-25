package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day25/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `        
	1=-0-2
	12111
	2=0=
	21
	2=01
	111
	20012
	112
	1=-1=
	1-12
	12
	1=
	122
	`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}

	solution := PartA(input)

	correct_answer := utilities.NewNumber("2=-1=0", 4890)

	if correct_answer.GetValue() != 4890 {
		t.Fatal("Expected correct_answer to have the value", 4890, "but it had", correct_answer.GetValue())
	}

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestConv(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 2022, 12345, 314159265,
		1747, 906, 198, 11, 201, 31, 1257, 32, 353, 107, 7, 3, 37, 4890}
	reprs := []string{"1", "2", "1=", "1-", "10", "11", "12", "2=", "2-", "20", "1=0",
		"1-0", "1=11-2", "1-0---0", "1121-1110-1=0", "1=-0-2", "12111", "2=0=", "21",
		"2=01", "111", "20012", "112", "1=-1=", "1-12", "12", "1=", "122", "2=-1=0"}

	for i, repr := range reprs {
		value := values[i]

		num := utilities.NewNumber(repr, 0)

		if num.GetValue() != value {
			t.Fatal("Expected", repr, "to map to", value, "but it mapped to", num.GetValue())
		}
	}

	for i, value := range values {
		num := utilities.NewNumber("", value)

		repr := reprs[i]

		if num.GetRepr() != repr {
			t.Fatal("Expected", value, "to map to", repr, "but it mapped to", num.GetRepr())
		}
	}

}
