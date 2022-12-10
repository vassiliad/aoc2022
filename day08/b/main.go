package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day08/a/utilities"
)

func ScenicScoreOneTree(world *utilities.World, x, y int) int {
	if x == 0 || x == world.Width-1 || y == 0 || y == world.Height-1 {
		return 0
	}

	s_up, s_left, s_down, s_right := 0, 0, 0, 0
	me := world.Trees[y*world.Width+x]

	for dy := 1; y+dy < world.Height; dy++ {
		s_down++

		if world.Trees[(y+dy)*world.Width+x] >= me {
			break
		}
	}

	for dy := 1; y >= dy; dy++ {
		other := world.Trees[(y-dy)*world.Width+x]
		s_up++

		if other >= me {
			break
		}
	}

	for dx := 1; x+dx < world.Width; dx++ {
		other := world.Trees[y*world.Width+x+dx]
		s_right++
		if other >= me {
			break
		}
	}

	for dx := 1; x >= dx; dx++ {
		other := world.Trees[y*world.Width+x-dx]
		s_left++
		if other >= me {
			break
		}
	}

	score := s_up * s_down * s_left * s_right
	return score
}

func PartB(world *utilities.World) int {
	best_score := 0
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			this_tree := ScenicScoreOneTree(world, x, y)

			if this_tree > best_score {
				best_score = this_tree
			}
		}
	}

	return best_score
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
