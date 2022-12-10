package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	rations []uint64
}

func (e *Elf) Calories() uint64 {
	var sum uint64 = 0

	for _, ration := range e.rations {
		sum += ration
	}

	return sum
}

func ReadScanner(scanner *bufio.Scanner) ([]Elf, error) {
	elves := []Elf{}

	curr_elf := Elf{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			if len(curr_elf.rations) == 0 {
				continue
			}

			elves = append(elves, curr_elf)
			curr_elf = Elf{}
			continue
		}
		number, err := strconv.ParseUint(line, 10, 0)

		if err != nil {
			return elves, err
		}
		curr_elf.rations = append(curr_elf.rations, number)
	}

	if len(curr_elf.rations) > 0 {
		elves = append(elves, curr_elf)
	}

	return elves, scanner.Err()
}

func ReadString(text string) ([]Elf, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Elf, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
