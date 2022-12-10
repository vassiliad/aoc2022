package main

import (
	"os"
	"strings"

	utilities "github.com/vassiliad/aoc2022/day03/a/utilities"
)

func PartB(rucksacks []utilities.Rucksack) uint64 {
	var score uint64 = 0

	for i := 0; i < len(rucksacks); i += 3 {
		sacks := []string{
			rucksacks[i].Combined(),
			rucksacks[i+1].Combined(),
			rucksacks[i+2].Combined(),
		}

		for _, letter := range sacks[0] {
			if strings.ContainsRune(sacks[1], letter) && strings.ContainsRune(sacks[2], letter) {
				score += (*utilities.Rucksack)(nil).PriorityItem(letter)
				break
			}
		}
	}

	return score
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
		os.Exit(1)
	}

	logger.Println("Read", len(input), "items")
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
