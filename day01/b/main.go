package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day01/a/utilities"
)

/*Identify the 3 elves with the most calories and return their combined calories*/
func PartB(elves []utilities.Elf) uint64 {
	find_largest_xth := func(pos int) uint64 {
		largest_idx := pos
		largest_cals := elves[pos].Calories()

		for idx, elf := range elves[pos+1:] {
			cals := elf.Calories()
			if cals > largest_cals {
				largest_cals = cals
				largest_idx = idx + pos + 1
			}
		}

		elves[pos], elves[largest_idx] = elves[largest_idx], elves[pos]

		return largest_cals
	}

	var sum uint64 = 0

	for i := 0; i < 3; i++ {
		sum += find_largest_xth(i)
	}

	return sum
}

func main() {
	logger := utilities.SetupLogger()

	elves, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(elves), "elves")
	sol := PartB(elves)
	logger.Println("Solution is", sol)
}
