package main

import (
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day20/a/utilities"
)

type Number struct {
	Value int
	Index int
}

type Metadata struct {
	Order      []Number
	MixedOrder []int
	Zero       *Number
}

func (m *Metadata) ToSlice() []int {
	ret := make([]int, len(m.MixedOrder))

	zero_index := m.Zero.Index

	for i := 0; i < len(m.Order); i++ {
		ret[i] = m.Order[m.MixedOrder[(zero_index+i)%len(m.MixedOrder)]].Value
	}

	return ret
}

func (m *Metadata) ToIndexSlice() []int {
	ret := make([]int, len(m.MixedOrder))

	for i, idx := range m.MixedOrder {
		ret[i] = m.Order[idx].Index
	}

	return ret
}

func (m *Metadata) Report() {
	fmt.Printf("%+v\n", m.ToSlice())
}

func (m *Metadata) swap(left, right int) {
	// VV: There can be super small and super large numbers.
	// Make them fit inside [0, len(numbers)-1]
	size := len(m.MixedOrder)
	times_wrap_left := utilities.AbsInt(left/size) + 1
	times_wrap_right := utilities.AbsInt(right/size) + 1

	left = (times_wrap_left*size + left) % size
	right = (times_wrap_right*size + right) % size

	idx_l := m.MixedOrder[left]
	idx_r := m.MixedOrder[right]

	m.MixedOrder[left], m.MixedOrder[right] = idx_r, idx_l
	m.Order[idx_l].Index, m.Order[idx_r].Index = m.Order[idx_r].Index, m.Order[idx_l].Index
}

func (m *Metadata) MoveForward(idx, spaces int) {
	mixed_index := m.Order[idx].Index

	for i := mixed_index + 1; i <= mixed_index+spaces; i++ {
		m.swap(i-1, i)
	}
}

func (m *Metadata) MoveBackward(idx, spaces int) {
	mixed_index := m.Order[idx].Index

	for i := mixed_index; i > mixed_index-spaces; i-- {
		m.swap(i-1, i)
	}
}

func (m *Metadata) DecideSpaces(idx int) int {
	number := &m.Order[idx]
	spaces := number.Value

	if utilities.AbsInt(spaces) >= len(m.Order) {
		// VV: when looping around "this number" won't exist.
		// There will be "len(numbers) -1" numbers
		size := len(m.Order) - 1
		dest_index := (spaces + number.Index) % size
		spaces = dest_index - number.Index
	}

	return spaces
}

func (m *Metadata) MoveNumber(idx int) {
	spaces := m.DecideSpaces(idx)

	// VV: There's no need to move anything if we're moving exactly len(numbers) spaces
	// because we'd end up moving around the entire list
	if spaces == 0 {
		return
	}

	if spaces > 0 {
		m.MoveForward(idx, spaces)
	} else {
		m.MoveBackward(idx, -spaces)
	}
}

func (m *Metadata) Prepare(numbers []int, key int) {
	order := make([]Number, len(numbers))
	mixed_order := make([]int, len(numbers))

	index_value_zero := -1

	for index, value := range numbers {
		order[index].Index = index
		order[index].Value = value * key
		mixed_order[index] = index

		if value == 0 {
			index_value_zero = index
		}
	}

	m.MixedOrder = mixed_order
	m.Order = order
	m.Zero = &order[index_value_zero]
}

func (m *Metadata) Decrypt() {
	for i := 0; i < 10; i++ {
		fmt.Println("Decrypting", i+1, "out of", 10)
		for idx := 0; idx < len(m.Order); idx++ {
			m.MoveNumber(idx)
		}
		if len(m.MixedOrder) < 40 {
			m.Report()
		}
	}
}

func (m *Metadata) ToGroveCoords() [3]int {
	grove := [3]int{
		m.Order[m.MixedOrder[(m.Zero.Index+1000)%len(m.Order)]].Value,
		m.Order[m.MixedOrder[(m.Zero.Index+2000)%len(m.Order)]].Value,
		m.Order[m.MixedOrder[(m.Zero.Index+3000)%len(m.Order)]].Value,
	}
	return grove
}

func PartB(numbers []int) (Metadata, int) {
	const key = 811589153

	metadata := Metadata{}
	metadata.Prepare(numbers, key)

	fmt.Println("Zero starts at", metadata.Zero.Index)

	metadata.Decrypt()

	fmt.Println("Zero ends up at", metadata.Zero.Index)

	grove := metadata.ToGroveCoords()
	return metadata, grove[0] + grove[1] + grove[2]
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	metadata, sol := PartB(input)
	logger.Println("Solution is", sol, "Grove coordinates are", metadata.ToGroveCoords())
}
