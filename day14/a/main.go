package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day14/a/utilities"
)

func SimulateSand(world *utilities.World) bool {
	x, y := 500, 0

	for {
		if y > world.Bottom {
			return false
		}

		if world.TileGet(x, y+1) == utilities.TileEmpty {
			y++
		} else if world.TileGet(x-1, y+1) == utilities.TileEmpty {
			y++
			x--
		} else if world.TileGet(x+1, y+1) == utilities.TileEmpty {
			y++
			x++
		} else {
			break
		}
	}

	world.TilePut(x, y, utilities.TileStationarySand)
	return true
}

func PartA(world *utilities.World) int {
	grains := 0

	for SimulateSand(world) {
		grains++
	}

	return grains
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
