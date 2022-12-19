package main

import (
	"testing"

	"github.com/vassiliad/aoc2022/day19/a/utilities"
)

func TestSmall(t *testing.T) {
	small := `
	Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
    Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
    `

	input, err := utilities.ReadString(small)

	ore := utilities.Recipe{
		Output:       utilities.ResOre,
		CostOre:      4,
		CostClay:     0,
		CostObsidian: 0,
	}

	clay := utilities.Recipe{
		Output:       utilities.ResClay,
		CostOre:      2,
		CostClay:     0,
		CostObsidian: 0,
	}

	obsidian := utilities.Recipe{
		Output:       utilities.ResObsidian,
		CostOre:      3,
		CostClay:     14,
		CostObsidian: 0,
	}

	geode := utilities.Recipe{
		Output:       utilities.ResGeode,
		CostOre:      2,
		CostClay:     0,
		CostObsidian: 7,
	}

	if ore != input[0].Robots[utilities.ResOre] {
		t.Fatalf("%+v\n", input[1].Robots[utilities.ResOre])
	}

	if clay != input[0].Robots[utilities.ResClay] {
		t.Fatalf("%+v\n", input[1].Robots[utilities.ResClay])
	}

	if obsidian != input[0].Robots[utilities.ResObsidian] {
		t.Fatalf("%+v\n", input[1].Robots[utilities.ResObsidian])
	}

	if geode != input[0].Robots[utilities.ResGeode] {
		t.Fatalf("%+v\n", input[1].Robots[utilities.ResGeode])
	}

	if err != nil {
		t.Fatal("Run into problems while reading input. Problem", err)
	}
	solution := PartA(input)

	const correct_answer = 33

	if solution != correct_answer {
		t.Fatal("Expected answer to be", correct_answer, "but it was", solution)
	}
}
