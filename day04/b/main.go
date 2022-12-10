package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day04/a/utilities"
)

type SmarterSections struct {
	utilities.Sections
}

func Overlap(first, second *utilities.Sections) bool {
	/*VV:
	|---One--------|
	        -delta-
	       |---Another---|
	if delta = another.Begin - one.End <= 0 then there's an overlap
	*/

	return (first.End >= second.End && first.Begin <= second.End) ||
		(second.End >= first.End && second.Begin <= first.End)
}

func PartB(assignments []utilities.Assignment) uint64 {
	var num_pairs uint64 = 0

	for _, assignment := range assignments {
		if Overlap(&assignment.First, &assignment.Second) {
			num_pairs += 1
		}
	}

	return num_pairs
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(input), "items")
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
