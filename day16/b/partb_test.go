package main

import (
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
	solution := PartB(input)

	const correct_answer = 1707

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}

// func TestDummy(t *testing.T) {
// 	small := `
// 	Valve AA has flow rate=0; tunnels lead to valves BB
// 	Valve BB has flow rate=0; tunnels lead to valves CC, AA
// 	Valve CC has flow rate=0; tunnels lead to valves BB, FF, DD, EE
// 	Valve FF has flow rate=0; tunnels lead to valves GG, CC
// 	Valve GG has flow rate=100000; tunnels lead to valves FF
// 	Valve EE has flow rate=0; tunnels lead to valves CC, HH
// 	Valve HH has flow rate=2; tunnels lead to valves EE
// 	Valve DD has flow rate=0; tunnels lead to valves CC, II
// 	Valve II has flow rate=0; tunnels lead to valves DD, KK
// 	Valve KK has flow rate=1000; tunnels lead to valves II
// 	`

// 	input, err := utilities.ReadString(small)

// 	if err != nil {
// 		t.Fatal("Run into problems while reading input. Problem", err)
// 	}
// 	solution := PartB(input)

// 	gg := 100000
// 	kk := 1000
// 	hh := 1

// 	correct_answer := gg*(26-6) + kk*(26-7) + hh*(26-11)

// 	if solution != correct_answer {
// 		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
// 	}
// }
