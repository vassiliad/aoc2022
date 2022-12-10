package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DecisionOpponent byte

type DecisionMine byte

type Round struct {
	other DecisionOpponent
	mine  DecisionMine
	score *int
}

const (
	OpponentRock     DecisionOpponent = 'A'
	OpponentPaper    DecisionOpponent = 'B'
	OpponentScissors DecisionOpponent = 'C'
	MineRock         DecisionMine     = 'X'
	MinePaper        DecisionMine     = 'Y'
	MineScissors     DecisionMine     = 'Z'
)

func (r *Round) Score() int {
	if r.score != nil {
		return *r.score
	}

	compute := func() int {
		sum := 0

		if r.mine == MineRock {
			sum = 1
			if r.other == OpponentScissors {
				sum += 6
			} else if r.other == OpponentRock {
				sum += 3
			}
		} else if r.mine == MinePaper {
			sum = 2
			if r.other == OpponentRock {
				sum += 6
			} else if r.other == OpponentPaper {
				sum += 3
			}
		} else {
			// This is MineScissors
			sum = 3
			if r.other == OpponentPaper {
				sum += 6
			} else if r.other == OpponentScissors {
				sum += 3
			}
		}

		return sum
	}

	r.score = new(int)
	*r.score = compute()

	return *r.score
}

func ReadScanner(scanner *bufio.Scanner) ([]Round, error) {
	rounds := []Round{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")
		round := Round{other: DecisionOpponent(tokens[0][0]), mine: DecisionMine(tokens[1][0])}

		if round.other != OpponentRock && round.other != OpponentPaper && round.other != OpponentScissors {
			return rounds, fmt.Errorf("OpponentHand %s is invalid", tokens[0])
		}

		if round.mine != MineRock && round.mine != MinePaper && round.mine != MineScissors {
			return rounds, fmt.Errorf("MyHand %s is invalid", tokens[1])
		}

		_ = round.Score()

		rounds = append(rounds, round)
	}

	return rounds, scanner.Err()
}

func ReadString(text string) ([]Round, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Round, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
