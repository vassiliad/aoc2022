package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	X int
	Y int
	Z int
}

func ReadScanner(scanner *bufio.Scanner) ([]Cube, error) {
	ret := []Cube{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, ",")

		cube := Cube{}

		var err error = nil

		cube.X, err = strconv.Atoi(tokens[0])

		if err != nil {
			return ret, fmt.Errorf("unable to parse X in %s due to %s", line, err)
		}

		cube.Y, err = strconv.Atoi(tokens[1])

		if err != nil {
			return ret, fmt.Errorf("unable to parse Y in %s due to %s", line, err)
		}

		cube.Z, err = strconv.Atoi(tokens[2])

		if err != nil {
			return ret, fmt.Errorf("unable to parse Z in %s due to %s", line, err)
		}

		ret = append(ret, cube)
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Cube, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Cube, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
