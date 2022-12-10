package utilities

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	From   uint
	To     uint
	Number uint
}

type Platform struct {
	// VV: The top crate is at position [0]
	Stacks []list.List
	Orders []Order
}

func (p *Platform) TopCrates() string {
	ret := []byte{}

	for _, stack := range p.Stacks {
		if stack.Front() != nil {
			val := stack.Front().Value
			ret = append(ret, (val).(byte))
		} else {
			ret = append(ret, ' ')
		}
	}

	return string(ret)
}

func (p *Platform) StackString(idx int) string {
	ret := []byte{}

	next := p.Stacks[idx].Front()
	for next != nil {
		ret = append(ret, next.Value.(byte))
		next = next.Next()
	}

	return string(ret)
}

func (p *Platform) Move(order *Order, just_one_crate bool) {
	src := &p.Stacks[order.From]
	dst := &p.Stacks[order.To]
	total_moved := 0

	for total_moved < int(order.Number) {
		to_move := 1

		if !just_one_crate {
			to_move = int(order.Number)
		}

		if src.Len() < to_move {
			to_move = src.Len()
		}

		if to_move == 0 {
			panic(fmt.Sprintf("order: %v, stacks: %v", order, p.Stacks))
		}

		ptr := src.Front()

		for i := 0; i < to_move-1; i++ {
			ptr = ptr.Next()
		}

		for i := 0; i < to_move; i++ {
			prev := ptr.Prev()
			val := src.Remove(ptr)
			dst.PushFront(val)
			ptr = prev
		}

		total_moved += to_move
	}

}

func ReadScanner(scanner *bufio.Scanner) (*Platform, error) {
	platform := new(Platform)
	parsed_crates := false

	// VV: Parse the crates first, if you encounter an empty line
	// after you've started parsing then move over to parsing
	// orders
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if !parsed_crates {
				continue
			}
			break
		}

		parsed_crates = true

		if len(platform.Stacks) == 0 {
			platform.Stacks = make([]list.List, (len(line)+1)/4)
		}

		for i := 0; i < len(line); i += 4 {
			if line[i] == '[' {
				platform.Stacks[i/4].PushBack(line[i+1])
			}
		}
	}

	if scanner.Err() != nil {
		return platform, scanner.Err()
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		order := Order{}

		if t, err := strconv.ParseUint(tokens[1], 10, 0); err != nil {
			return platform, err
		} else {
			order.Number = uint(t)
		}

		if t, err := strconv.ParseUint(tokens[3], 10, 0); err != nil {
			return platform, err
		} else {
			order.From = uint(t) - 1
		}

		if t, err := strconv.ParseUint(tokens[5], 10, 0); err != nil {
			return platform, err
		} else {
			order.To = uint(t) - 1
		}

		platform.Orders = append(platform.Orders, order)
	}

	return platform, scanner.Err()
}

func ReadString(text string) (*Platform, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Platform, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
