package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type OpCodeAdd struct {
	Cycle     int
	Register  string
	Immediate int
}

type OpCodeNoop struct {
}

func ReadScanner(scanner *bufio.Scanner) ([]any, error) {
	ret := []any{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "noop" {
			ret = append(ret, OpCodeNoop{})
		} else if tokens[0] == "addx" {
			steps, err := strconv.Atoi(tokens[1])

			if err != nil {
				return ret, err
			}

			ret = append(ret, OpCodeAdd{Register: "X", Immediate: steps})

		} else {
			logger := SetupLogger()
			logger.Panicf(line, tokens)
		}

	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]any, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]any, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
