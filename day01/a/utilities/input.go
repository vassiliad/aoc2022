package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	calories *uint64
}

func (e *Elf) Calories() uint64 {
	return *e.calories
}

func ReadScanner(scanner *bufio.Scanner) ([]Elf, error) {
	elves := []Elf{}

	curr_elf := Elf{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			if curr_elf.calories == nil {
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

		if curr_elf.calories == nil {
			curr_elf.calories = new(uint64)
			*curr_elf.calories = 0
		}
		*curr_elf.calories += number
	}

	if curr_elf.calories != nil {
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
