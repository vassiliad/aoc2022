package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day08/a/utilities"
)

func PartA(world *utilities.World) int {
	visible := make(map[int]int8)

	max := make([]int8, world.Width)
	for dy := 0; dy < world.Height; dy++ {
		max_horiz := int8(0)

		for dx := 0; dx < world.Width; dx++ {
			idx := dy*world.Width + dx

			if dy == 0 || (max[dx] < world.Trees[idx]) {
				max[dx] = world.Trees[idx]
				visible[idx] = 1
			}

			if dx == 0 || max_horiz < world.Trees[idx] {
				max_horiz = world.Trees[idx]
				visible[idx] = 1
			}
		}
	}

	for i := 0; i < world.Height; i++ {
		max[i] = 0
	}

	for dy := world.Height - 1; dy >= 1; dy-- {
		max_horiz := int8(0)
		for dx := world.Width - 1; dx > -1; dx-- {
			idx := dy*world.Width + dx

			if (dy == world.Height-1) || (max[dx] < world.Trees[idx]) {
				max[dx] = world.Trees[idx]
				visible[idx] = 1
			}

			if (dx == world.Width-1) || (max_horiz < world.Trees[idx]) {
				max_horiz = world.Trees[idx]
				visible[idx] = 1
			}
		}
	}

	return len(visible)
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
