package utilities

import (
	"bufio"
	"os"
	"strings"
)

type World struct {
	Trees  []int8
	Width  int
	Height int
}

func ReadScanner(scanner *bufio.Scanner) (*World, error) {
	world := new(World)

	curr_line := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		world.Height = len(line)
		world.Width = len(line)

		if len(world.Trees) == 0 {
			world.Trees = make([]int8, world.Width*world.Height)
		}

		for i, letter := range line {
			height := letter - '0'
			world.Trees[curr_line*world.Width+i] = int8(height)
		}

		curr_line += 1
	}

	return world, scanner.Err()
}

func ReadString(text string) (*World, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*World, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
