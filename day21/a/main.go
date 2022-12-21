package main

import (
	"fmt"
	"os"

	"github.com/vassiliad/aoc2022/day21/a/utilities"
)

func Compute(calc *utilities.Calculator, res string) int {
	// fmt.Printf("%+v\n", *calc)

	remaining := []string{res}
	for len(remaining) > 0 {
		res = remaining[len(remaining)-1]

		if _, ok := calc.Registers[res]; ok {
			remaining = remaining[:len(remaining)-1]
			continue
		} else {
			// VV: We have not computed the value YET. If we know the values of operands, compute the value now
			// else, just ask future loops to compute the operands for us
			expr, ok := calc.Expressions[res]

			if !ok {
				panic(fmt.Sprintf("Could not find expression for %s", res))
			}

			missing := 0
			to_add := []string{}
			if _, ok = calc.Registers[expr.Left]; !ok {
				missing++
				to_add = append(to_add, expr.Left)
			}

			if _, ok = calc.Registers[expr.Right]; !ok {
				missing++
				to_add = append(to_add, expr.Right)
			}

			if missing == 0 {
				value, _ := calc.Compute(&expr)
				calc.Registers[res] = value
				remaining = remaining[0 : len(remaining)-1]
			} else {
				remaining = append(remaining, to_add...)
			}
		}
	}

	return calc.Registers[res]
}

func PartA(calculator *utilities.Calculator) int {
	return Compute(calculator, "root")
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartA(input)
	logger.Println("Solution is", sol)
}
