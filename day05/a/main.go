package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day05/a/utilities"
)

func PartA(platform *utilities.Platform) string {
	for _, order := range platform.Orders {
		platform.Move(&order, true)
	}

	return platform.TopCrates()
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(input.Orders), "items")
	sol := PartA(input)
	logger.Println("Solution is", sol)
}
