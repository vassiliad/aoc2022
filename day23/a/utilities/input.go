package utilities

import (
	"bufio"
	"os"
	"strings"
)

type Direction int8
type Tile int8

const (
	DirN          Direction = 0
	DirS          Direction = 1
	DirW          Direction = 2
	DirE          Direction = 3
	TileContested           = -1
)

type Point struct {
	X int
	Y int
}

type Group struct {
	Left   int
	Right  int
	Top    int
	Bottom int
	Elves  map[Point]int8
}

func ReadScanner(scanner *bufio.Scanner) (*Group, error) {
	ret := Group{Elves: map[Point]int8{}}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		ret.Right = MaxInt(ret.Right, len(line))
		for x, c := range line {
			if c == '#' {
				ret.Elves[Point{X: x, Y: ret.Bottom}] = 1
			}
		}
		ret.Bottom--
	}

	return &ret, scanner.Err()
}

func ReadString(text string) (*Group, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Group, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
