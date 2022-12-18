package main

import (
	"container/list"
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day18/a/utilities"
)

func FloodFill(low, high *utilities.Cube, book map[utilities.Cube]int8) int {
	flooded := map[utilities.Cube]int8{}
	remaining := list.New()

	directions := []utilities.Cube{
		{X: -1}, {X: +1},
		{Y: -1}, {Y: +1},
		{Z: -1}, {Z: +1},
	}

	remaining.PushFront(*low)

	for remaining.Len() > 0 {
		first := remaining.Remove(remaining.Back()).(utilities.Cube)
		flooded[first] = 1

		for _, d := range directions {
			n := first
			n.X += d.X
			n.Y += d.Y
			n.Z += d.Z

			if n.X < low.X || n.Y < low.Y || n.Z < low.Z ||
				n.X > high.X || n.Y > high.Y || n.Z > high.Z ||
				flooded[n] == 1 || book[n] == 1 {
				continue
			}

			remaining.PushBack(n)
		}
	}

	exposed := 0

	for cube, _ := range book {
		sides := []utilities.Cube{
			{X: cube.X - 1, Y: cube.Y, Z: cube.Z},
			{X: cube.X + 1, Y: cube.Y, Z: cube.Z},
			{X: cube.X, Y: cube.Y - 1, Z: cube.Z},
			{X: cube.X, Y: cube.Y + 1, Z: cube.Z},
			{X: cube.X, Y: cube.Y, Z: cube.Z - 1},
			{X: cube.X, Y: cube.Y, Z: cube.Z + 1},
		}

		for _, neighbour := range sides {
			if flooded[neighbour] == 1 {
				exposed++
			}
		}
	}

	return exposed
}

func PartB(cubes []utilities.Cube) int {
	book := map[utilities.Cube]int8{}
	low, high := cubes[0], cubes[0]

	const extra = 1
	for _, cube := range cubes {
		low.X = utilities.MinInt(low.X, cube.X-extra)
		low.Y = utilities.MinInt(low.Y, cube.Y-extra)
		low.Z = utilities.MinInt(low.Z, cube.Z-extra)

		high.X = utilities.MaxInt(high.X, cube.X+extra)
		high.Y = utilities.MaxInt(high.Y, cube.Y+extra)
		high.Z = utilities.MaxInt(high.Z, cube.Z+extra)

		book[cube] = 1
	}

	fmt.Printf("Low %+v, High %+v\n", low, high)

	return FloodFill(&low, &high, book)
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
