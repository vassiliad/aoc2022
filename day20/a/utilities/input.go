package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadScanner(scanner *bufio.Scanner) ([]int, error) {
	ret := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		number, err := strconv.Atoi(line)

		if err != nil {
			return ret, fmt.Errorf("unable to parse %s due to %s", line, err)
		}

		ret = append(ret, number)
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]int, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
