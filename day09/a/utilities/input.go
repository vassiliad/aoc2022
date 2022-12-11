package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Direction int8

const (
	DirectionLeft  Direction = 0
	DirectionRight Direction = 1
	DirectionUp    Direction = 2
	DirectionDown  Direction = 3
)

type Move struct {
	Direction Direction
	Steps     int
}

func ReadScanner(scanner *bufio.Scanner) ([]Move, error) {
	ret := []Move{}

	directions := map[string]Direction{
		"L": DirectionLeft,
		"R": DirectionRight,
		"U": DirectionUp,
		"D": DirectionDown,
	}

	curr_line := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		steps, err := strconv.Atoi(tokens[1])

		if err != nil {
			return ret, err
		}

		ret = append(ret, Move{Direction: directions[tokens[0]], Steps: int(steps)})

		curr_line += 1
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Move, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Move, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
