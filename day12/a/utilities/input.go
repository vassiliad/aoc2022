package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

type World struct {
	Tiles  []int8
	Width  int
	Height int

	Start Point
	End   Point
}

func ReadScanner(scanner *bufio.Scanner) (*World, error) {
	ret := World{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		for x, val := range line {
			if val == 'S' {
				val = 'a'
				ret.Start.Y = ret.Height
				ret.Start.X = x
			} else if val == 'E' {
				val = 'z'
				ret.End.Y = ret.Height
				ret.End.X = x
			}

			ret.Tiles = append(ret.Tiles, int8(val-'a'))
		}

		ret.Width = len(line)
		ret.Height++
	}

	return &ret, scanner.Err()
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
