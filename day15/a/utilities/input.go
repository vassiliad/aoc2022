package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Pair struct {
	Sensor   Point
	Beacon   Point
	Distance int
}

func (p *Point) Distance(other *Point) int {
	return AbsInt(p.X-other.X) + AbsInt(p.Y-other.Y)
}

func ParsePoint(x, y string) (*Point, error) {
	n1, err := strconv.Atoi(x)

	if err != nil {
		return nil, err
	}

	n2, err := strconv.Atoi(y)

	if err == nil {
		return &Point{X: n1, Y: n2}, nil
	}

	return nil, err
}

func ReadScanner(scanner *bufio.Scanner) ([]Pair, error) {
	ret := []Pair{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		sensor, err := ParsePoint(tokens[2][2:len(tokens[2])-1], tokens[3][2:len(tokens[3])-1])

		if err != nil {
			return ret, fmt.Errorf("unable to decobe sensor in %s because %s", line, err)
		}

		beacon, err := ParsePoint(tokens[8][2:len(tokens[8])-1], tokens[9][2:len(tokens[9])])

		if err != nil {
			return ret, fmt.Errorf("unable to decobe beacon in %s because %s", line, err)
		}

		ret = append(ret, Pair{Beacon: *beacon, Sensor: *sensor, Distance: sensor.Distance(beacon)})
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Pair, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Pair, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
