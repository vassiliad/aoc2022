package main

import (
	"strconv"
	"testing"

	"github.com/vassiliad/aoc2022/day16/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	Valve BB has flow rate=13; tunnels lead to valves CC, AA
	Valve CC has flow rate=2; tunnels lead to valves DD, BB
	Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
	Valve EE has flow rate=3; tunnels lead to valves FF, DD
	Valve FF has flow rate=0; tunnels lead to valves EE, GG
	Valve GG has flow rate=0; tunnels lead to valves FF, HH
	Valve HH has flow rate=22; tunnel leads to valve GG
	Valve II has flow rate=0; tunnels lead to valves AA, JJ
	Valve JJ has flow rate=21; tunnel leads to valve II`

	input, err := utilities.ReadString(small)

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartA(input)

	const correct_answer = 1651

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

func TestBitMask(t *testing.T) {
	var bm utilities.BitMask = 0

	if bm.Get(33) != false {
		t.Fatal("Expected BitMask to be empty")
	}

	bm.Set(0)
	bm.Set(2)
	bm.Set(4)

	for i := 0; i <= 4; i++ {
		if bm.Get(i) != (i%2 == 0) {
			t.Fatalf("BitMask.%d=%t is unexpected, full bitmask is %s",
				i, bm.Get(i), strconv.FormatUint(uint64(bm), 2))
		}
	}

	bm.UnSet(4)

	if bm.Get(4) == true {
		t.Fatalf("Expected BitMask.4 == false - %s", strconv.FormatUint(uint64(bm), 2))
	}

	bm = 0
	for i := 0; i < 3; i++ {
		bm.Set(i)
	}

	var all_valves uint64 = 1<<uint64(3) - 1

	if bm != utilities.BitMask(all_valves) {
		t.Fatalf("Expected BitMask %s but was %s",
			strconv.FormatUint(uint64(all_valves), 2), strconv.FormatUint(uint64(bm), 2))
	}
}
