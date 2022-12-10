package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Sections struct {
	Begin uint
	End   uint
}

type Assignment struct {
	First  Sections
	Second Sections
}

func (s *Sections) SupersetOf(other *Sections) bool {
	return (s.Begin <= other.Begin) && (s.End >= other.End)
}

func (s *Sections) Parse(text string) error {
	halves := strings.Split(text, "-")
	begin, err := strconv.ParseUint(halves[0], 10, 0)

	if err != nil {
		return err
	}

	s.Begin = uint(begin)

	end, err := strconv.ParseUint(halves[1], 10, 0)

	if err != nil {
		return err
	}

	s.End = uint(end)
	return nil
}

func ReadScanner(scanner *bufio.Scanner) ([]Assignment, error) {
	assignments := []Assignment{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		assignment := Assignment{}
		halves := strings.Split(line, ",")

		if err := assignment.First.Parse(halves[0]); err != nil {
			return assignments, err
		}

		if err := assignment.Second.Parse(halves[1]); err != nil {
			return assignments, err
		}

		assignments = append(assignments, assignment)
	}

	return assignments, scanner.Err()
}

func ReadString(text string) ([]Assignment, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Assignment, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
