package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day07/a/utilities"
)

func PartA(root *utilities.Filesystem) uint64 {
	const threshold = uint64(100000)
	sum := uint64(0)
	for _, v := range *root {
		if v.CountSize(root) <= threshold {
			sum += v.CountSize(root)
		}
	}

	return sum
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	logger.Println("Read", len(input), "items")
	sol := PartA(&input)
	logger.Println("Solution is", sol)
}
