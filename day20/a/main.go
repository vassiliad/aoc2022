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

func (m *Metadata) Report(idx int) {
	fmt.Printf("%2d : ", m.Order[idx].Value)
	for i, idx := range m.MixedOrder {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%2d", m.Order[idx].Value)
	}
	fmt.Printf("\n")
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

	// VV: we are simulating removing the number, then going round and round a bunch of times in a circular buffer.
	// There will be "len(numbers) -1" numbers till we put the number back in.
	size := len(m.Order) - 1
	dest_index := (spaces + number.Index) % size
	spaces = dest_index - number.Index

	return spaces
}

func (m *Metadata) MoveNumber(idx int) {
	spaces := m.DecideSpaces(idx)

	if spaces > 0 {
		m.MoveForward(idx, spaces)
	} else {
		m.MoveBackward(idx, -spaces)
	}
}

func (m *Metadata) Prepare(numbers []int) {
	order := make([]Number, len(numbers))
	mixed_order := make([]int, len(numbers))

	index_value_zero := -1

	for index, value := range numbers {
		order[index].Index = index
		order[index].Value = value
		mixed_order[index] = index

		if value == 0 {
			index_value_zero = index
		}
	}

	m.MixedOrder = mixed_order
	m.Order = order
	m.Zero = &order[index_value_zero]
}

func (m *Metadata) Encrypt() {
	for idx := 0; idx < len(m.Order); idx++ {
		m.MoveNumber(idx)
		if len(m.MixedOrder) < 40 {
			m.Report(idx)
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

func PartA(numbers []int) (Metadata, int) {
	metadata := Metadata{}
	metadata.Prepare(numbers)

	fmt.Println("Zero starts at", metadata.Zero.Index)

	metadata.Encrypt()

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
	metadata, sol := PartA(input)
	logger.Println("Solution is", sol, "Grove coordinates are", metadata.ToGroveCoords())
}
