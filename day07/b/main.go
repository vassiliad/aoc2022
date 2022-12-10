package main

import (
	"math"
	"os"

	utilities "github.com/vassiliad/aoc2022/day07/a/utilities"
)

func PartB(root *utilities.Filesystem) uint64 {
	const threshold = uint64(30000000)
	const capacity = uint64(70000000)

	root_used := (*root)["/"].CountSize(root)

	min_to_delete := uint64(math.MaxUint64)

	for _, v := range *root {
		size := v.CountSize(root)

		if (capacity-root_used+size) >= threshold && min_to_delete > size {
			min_to_delete = size
		}
	}

	return min_to_delete
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	logger.Println("Read", len(input), "items")
	sol := PartB(&input)
	logger.Println("Solution is", sol)
}
