package main

import (
	"os"

	old_utils "github.com/vassiliad/aoc2022/day02/a/utilities"
	utilities "github.com/vassiliad/aoc2022/day02/b/utilities"
)

func PartB(rounds []utilities.Round) int {
	sum := 0

	for _, r := range rounds {
		sum += r.Score()
	}

	return sum
}

func main() {
	logger := old_utils.SetupLogger()

	logger.Println("Parse input")
	rounds, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(rounds), "elves")
	sol := PartB(rounds)
	logger.Println("Solution is", sol)
}
