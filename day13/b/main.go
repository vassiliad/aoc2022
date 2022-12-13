package main

import (
	"fmt"
	"os"
	"sort"

	utilities "github.com/vassiliad/aoc2022/day13/a/utilities"
)

const (
	EndLower  = -1
	Same      = 0
	EndHigher = 1
)

func MakeNewNested(old *utilities.Packet) *utilities.Packet {
	packet := utilities.Packet{}
	packet.Nested = append(packet.Nested, old)

	return &packet
}

func IsLower(left, right *utilities.Packet) int {
	if left.Value == nil && right.Value != nil {
		right = MakeNewNested(right)
	} else if left.Value != nil && right.Value == nil {
		left = MakeNewNested(left)
	}

	if left.Value != nil && right.Value != nil {
		return *(left.Value) - *(right.Value)
	} else if left.Value == nil && right.Value == nil {
		size_left := len(left.Nested)
		size_right := len(right.Nested)

		comparisons := Same

		for i := 0; i < utilities.MaxInt(size_left, size_right); i++ {
			if i == size_left {
				return EndLower
			} else if i == size_right {
				return EndHigher
			}

			comparisons |= IsLower(left.Nested[i], right.Nested[i])

			if comparisons != Same {
				return comparisons
			}
		}

		return comparisons
	}

	panic(fmt.Sprintf("Cannot compare %+v with %+v", left, right))
}

func PartB(packets []*utilities.Packet) int {
	two := 2
	six := 6

	// VV: I hate this, and that ^^
	div1 := utilities.Packet{Nested: []*utilities.Packet{{Nested: []*utilities.Packet{{Value: &two}}}}}
	div2 := utilities.Packet{Nested: []*utilities.Packet{{Nested: []*utilities.Packet{{Value: &six}}}}}

	packets = append(packets, &div1, &div2)

	sort.Slice(packets, func(i, j int) bool {
		return IsLower(packets[i], packets[j]) < Same
	})

	ret := 1

	for i, p := range packets {
		if p.Nested != nil && len(p.Nested) == 1 && (*p.Nested[0]).Value == nil && len(p.Nested[0].Nested) == 1 {
			uber_nested := p.Nested[0].Nested[0]

			if uber_nested.Value != nil && (*uber_nested.Value == 2 || *uber_nested.Value == 6) {
				ret *= i + 1
			}
		}
	}

	return ret
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
