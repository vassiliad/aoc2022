package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Number struct {
	repr  string
	value int
}

const (
	Zero     = '0'
	One      = '1'
	Two      = '2'
	MinusOne = '-'
	MinusTwo = '='
)

func (n *Number) GetRepr() string {
	return n.repr
}

func (n *Number) GetValue() int {
	return n.value
}

func (n *Number) SetRepr(repr string) {
	n.repr = repr
	n.recomputeValue()
}

func (n *Number) SetValue(value int) {
	n.value = value
	n.recomputeRepr()
}

func (n *Number) recomputeRepr() {
	n.repr = ""

	powers := []int{}
	value := n.value

	for value != 0 {
		digit := value % 5
		powers = append(powers, digit)
		value /= 5
	}

	for i := 0; i < len(powers); i++ {
		if powers[i] >= 3 {
			powers[i] -= 5

			if i == len(powers)-1 {
				powers = append(powers, 1)
			} else {
				powers[i+1]++
			}
		} else if powers[i] < -2 {
			panic(fmt.Sprintf("unexpected power p[%d]=%d", i, powers[i]))
		}
	}

	to_digit := map[int]byte{0: Zero, 1: One, 2: Two, -1: MinusOne, -2: MinusTwo}

	for i := len(powers) - 1; i > -1; i-- {
		if d, ok := to_digit[powers[i]]; ok {
			n.repr += string(d)
		} else {
			panic(fmt.Sprintf("unexpected power p[%d]=%d", i, powers[i]))
		}
	}
}

func (n *Number) recomputeValue() {
	power := 1
	n.value = 0

	for i := len(n.repr) - 1; i > -1; i-- {
		switch n.repr[i] {
		case One:
			n.value += power
		case Two:
			n.value += 2 * power
		case MinusOne:
			n.value -= power
		case MinusTwo:
			n.value -= 2 * power
		}

		power *= 5
	}
}

func NewNumber(repr string, value int) Number {
	number := Number{repr: repr, value: value}

	if len(repr) > 0 {
		number.recomputeValue()
	} else {
		number.recomputeRepr()
	}

	return number
}

func ReadScanner(scanner *bufio.Scanner) ([]Number, error) {
	ret := []Number{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		num := Number{repr: line}
		num.recomputeValue()

		ret = append(ret, num)
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Number, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Number, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
