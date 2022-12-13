package main

import (
	"fmt"
	"os"

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

func PartA(packets []*utilities.Packet) int {
	inorder := 0

	for i := 0; i < len(packets)/2; i++ {
		if IsLower(packets[i*2], packets[i*2+1]) < Same {
			inorder += i + 1
		}
	}

	return inorder
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartA(input)
	logger.Println("Solution is", sol)
}
