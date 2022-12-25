package main

import (
	"os"

	"github.com/vassiliad/aoc2022/day25/a/utilities"
)

func PartA(numbers []utilities.Number) utilities.Number {
	ret := utilities.Number{}

	for _, n := range numbers {
		sum := ret.GetValue() + n.GetValue()

		ret.SetValue(sum)
	}

	return ret
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartA(input)
	logger.Printf("Solution is %s = %d\n", sol.GetRepr(), sol.GetValue())
}
