package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Valley struct {
	Width     int
	Height    int
	Start     int
	End       int
	BlizLeft  map[int]map[int]int8
	BlizRight map[int]map[int]int8
	BlizUp    map[int]map[int]int8
	BlizDown  map[int]map[int]int8
}

func ReadScanner(scanner *bufio.Scanner) (*Valley, error) {
	ret := Valley{
		BlizLeft:  map[int]map[int]int8{},
		BlizRight: map[int]map[int]int8{},
		BlizUp:    map[int]map[int]int8{},
		BlizDown:  map[int]map[int]int8{},
		Start:     -1,
		End:       -1,
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if (line[1] == '#') || (line[len(line)-2] == '#') {
			dot := strings.Index(line, ".")
			if ret.Start == -1 {
				ret.Start = dot
				ret.Width = len(line)
				continue
			} else if ret.End == -1 {
				ret.End = dot
				continue
			} else {
				panic(fmt.Sprintf("Too many entry/exit points: %+v", ret))
			}
		}

		ret.BlizLeft[ret.Height] = map[int]int8{}
		ret.BlizRight[ret.Height] = map[int]int8{}
		ret.BlizUp[ret.Height] = map[int]int8{}
		ret.BlizDown[ret.Height] = map[int]int8{}

		for x, c := range line[:len(line)-1] {
			switch c {
			case '>':
				ret.BlizRight[ret.Height][x] = 1
			case '<':
				ret.BlizLeft[ret.Height][x] = 1
			case '^':
				ret.BlizUp[ret.Height][x] = 1
			case 'v':
				ret.BlizDown[ret.Height][x] = 1
			}
		}

		ret.Height++
	}

	return &ret, scanner.Err()
}

func ReadString(text string) (*Valley, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Valley, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
