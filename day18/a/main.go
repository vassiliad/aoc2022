package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day18/a/utilities"
)

func ExposedCubeSides(cube *utilities.Cube, book map[utilities.Cube]int8) int {
	sides := []utilities.Cube{
		{X: cube.X - 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X + 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X, Y: cube.Y - 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y + 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y, Z: cube.Z - 1},
		{X: cube.X, Y: cube.Y, Z: cube.Z + 1},
	}

	exposed_sides := 0
	for _, neighbour := range sides {
		if _, ok := book[neighbour]; !ok {
			exposed_sides++
		}
	}

	return exposed_sides
}

func PartA(cubes []utilities.Cube) int {
	book := map[utilities.Cube]int8{}

	for _, cube := range cubes {
		book[cube] = 1
	}

	exposed_sides := 0

	for _, cube := range cubes {
		exposed_sides += ExposedCubeSides(&cube, book)
	}

	return exposed_sides
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
