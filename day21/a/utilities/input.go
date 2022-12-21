package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator int8

const (
	OpMult Operator = 0
	OpDiv  Operator = 1
	OpAdd  Operator = 2
	OpSub  Operator = 3
)

type Expression struct {
	// VV: I don't need this for PartA, but might make my life easier for PartB
	Out   string
	Left  string
	Right string
	Op    Operator
}

type Calculator struct {
	Expressions map[string]Expression
	Registers   map[string]int
}

func (c *Calculator) Compute(expr *Expression) (int, bool) {
	missing := 0
	ok := false
	t1 := 0
	t2 := 0

	if t1, ok = c.Registers[expr.Left]; !ok {
		missing++
	}

	if t2, ok = c.Registers[expr.Right]; !ok {
		missing++
	}

	if missing == 0 {
		if expr.Op == OpAdd {
			return t1 + t2, true
		} else if expr.Op == OpSub {
			return t1 - t2, true
		} else if expr.Op == OpMult {
			return t1 * t2, true
		} else if expr.Op == OpDiv {
			return t1 / t2, true
		}
		panic(fmt.Sprintf("Invalid expression %+v", expr))
	} else {
		return 0, false
	}
}

func ReadScanner(scanner *bufio.Scanner) (*Calculator, error) {
	ret := Calculator{
		Expressions: map[string]Expression{},
		Registers:   map[string]int{},
	}

	op_map := map[string]Operator{
		"+": OpAdd,
		"-": OpSub,
		"*": OpMult,
		"/": OpDiv,
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")

		tokens := strings.Split(parts[1], " ")
		res := parts[0]

		if len(tokens) == 1 {
			num, err := strconv.Atoi(tokens[0])

			if err != nil {
				return &ret, fmt.Errorf("unable to decode MonkeyNumber %s due to %s", tokens[0], err)
			}

			ret.Registers[res] = num
		} else if len(tokens) == 3 {
			op := Operator(-1)
			if t, ok := op_map[tokens[1]]; ok {
				op = t
			}
			if op == -1 {
				return &ret, fmt.Errorf("unable to parse operator in %+v", tokens)
			}

			expr := Expression{Left: tokens[0], Right: tokens[2], Op: op, Out: res}
			if len(expr.Left) != 4 || len(expr.Right) != 4 || len(expr.Out) != 4 {
				return &ret, fmt.Errorf("invalid expression operands %+v", expr)
			}

			ret.Expressions[res] = expr
		} else {
			return &ret, fmt.Errorf("unexpected tokens %+v in line %s", tokens, line)
		}

	}

	return &ret, scanner.Err()
}

func ReadString(text string) (*Calculator, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Calculator, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
