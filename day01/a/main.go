package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day01/a/utilities"
)

func PartA(elves []utilities.Elf) uint64 {
	var max_cals uint64 = 0

	for _, elf := range elves {
		cals := elf.Calories()

		if cals > max_cals {
			max_cals = cals
		}
	}

	return max_cals
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	elves, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(elves), "elves")
	sol := PartA(elves)
	logger.Println("Solution is", sol)
}
