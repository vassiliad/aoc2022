package main

import (
	"io/ioutil"
	"os"

	utilities "github.com/vassiliad/aoc2022/day05/a/utilities"
)

func PartA(text string) int {
	const packet_size = 4

	for i := packet_size; i < len(text); i++ {
		all_unique := true
		for a := 0; a < packet_size-1 && all_unique; a++ {
			for b := a + 1; b < packet_size && all_unique; b++ {
				if text[i+a-packet_size] == text[i+b-packet_size] {
					all_unique = false
					break
				}
			}
		}

		if all_unique {
			return i
		}
	}

	return -1
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	f, err := os.Open(os.Args[1])
	if err != nil {
		logger.Fatalln("Run into problems while opening input. Problem", err)
	}

	input, err := ioutil.ReadAll(f)

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}

	text := string(input)

	logger.Println("Read", len(input), "items")
	sol := PartA(text)
	logger.Println("Solution is", sol)
}
