package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day04/a/utilities"
)

func PartA(assignments []utilities.Assignment) uint64 {
	var num_pairs uint64 = 0

	for _, assignment := range assignments {
		if assignment.First.SupersetOf(&assignment.Second) || assignment.Second.SupersetOf(&assignment.First) {
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
	sol := PartA(input)
	logger.Println("Solution is", sol)
}
