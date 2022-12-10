package utilities

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	first  string
	second string
}

func (r *Rucksack) Combined() string {
    return r.first + r.second
}

func (r *Rucksack) PriorityItem(letter rune) uint64 {
	chars := []byte(string(letter))

	if len(chars) != 1 {
		log.Panicln("Expected exactly 1 byte for rune", letter, "but got", chars)
	}

	this := chars[0]

	if 'a' <= this && this <= 'z' {
		return uint64(this) - uint64('a') + 1
	} else {
		return uint64(this) - uint64('A') + 27
	}
}

func (r *Rucksack) Score() uint64 {
	for _, letter := range r.first {
		if strings.ContainsRune(r.second, letter) {
			return r.PriorityItem(letter)
		}
	}

	// VV: Unreachable code
	panic(r)
}

func ReadScanner(scanner *bufio.Scanner) ([]Rucksack, error) {
	rucksacks := []Rucksack{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		rucksack := Rucksack{first: string(line[:len(line)/2]), second: string(line[len(line)/2:])}

		if len(rucksack.first) != len(rucksack.second) || len(rucksack.first) == 0 {
			panic(rucksack)
		}

		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks, scanner.Err()
}

func ReadString(text string) ([]Rucksack, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Rucksack, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
