package main

import (
	"fmt"
	"os"

	"github.com/vassiliad/aoc2022/day19/b/utilities"
)

type State struct {
	Items   utilities.ManyParts
	Robots  utilities.ManyParts
	Minutes int
}

func MaxGeodes(blueprint utilities.Blueprint, max_minutes int) int {
	remaining := []State{}
	max_required := utilities.ManyParts{}

	for _, r := range blueprint.Robots {
		max_required[utilities.ResOre] = utilities.MaxInt(max_required[utilities.ResOre], r.CostOre)
		max_required[utilities.ResClay] = utilities.MaxInt(max_required[utilities.ResClay], r.CostClay)
		max_required[utilities.ResObsidian] = utilities.MaxInt(max_required[utilities.ResObsidian], r.CostObsidian)
	}
	max_geodes := 0

	mayPush := func(state *State) {

		remaining = append(remaining, *state)
	}

	buildRobot := func(state State, robot_kind utilities.Resource) {
		robot := blueprint.Robots[robot_kind]
		// VV: I am not convinced that this check is correct - simply because it didn't work before
		// I added that +10 sign (it was trimming the path that lead to optimal results).
		// The idea here is that if you're already building as much of a resource in 1 minute
		// as the cost to build the most expensive robot (wrt to that resource) then you don't need
		// to build any other robots for this resource. That seems logical to me, yet here we are.
		// I even rewrote the whole thing from scratch thinking that I must have a logical error
		// elsewhere ...
		if robot_kind != utilities.ResGeode && state.Items[robot_kind] > max_required[robot_kind]+10 {
			return
		}
		robot.Pay(&state.Items)
		for i := 0; i < utilities.TotalResources; i++ {
			state.Items[i] += state.Robots[i]
		}
		state.Robots[robot_kind] += 1

		if max_geodes < state.Items[utilities.ResGeode] {
			max_geodes = state.Items[utilities.ResGeode]
			fmt.Printf("%d] %+v\n", blueprint.ID, state)
		}
		mayPush(&state)
	}

	initial := State{
		Robots:  utilities.ManyParts{1, 0, 0},
		Minutes: 0,
	}

	mayPush(&initial)

	visited := map[State]int8{}

	for len(remaining) > 0 {
		state := remaining[len(remaining)-1]
		remaining = remaining[:len(remaining)-1]

		mins_left := max_minutes - state.Minutes
		max_geode_bots := mins_left + state.Robots[utilities.ResGeode]
		theoretical_geodes := state.Items[utilities.ResGeode] + (mins_left * (mins_left + 1) / 2) + max_geode_bots

		if theoretical_geodes < max_geodes {
			continue
		}

		if visited[state] == 0 {
			visited[state] = 1
		} else {
			continue
		}

		state.Minutes++
		stockpile_res := state
		for i := 0; i < utilities.TotalResources; i++ {
			stockpile_res.Items[i] += stockpile_res.Robots[i]
		}

		if max_geodes < stockpile_res.Items[utilities.ResGeode] {
			max_geodes = stockpile_res.Items[utilities.ResGeode]
			fmt.Printf("%d] %+v\n", blueprint.ID, stockpile_res)
		}

		if state.Minutes == max_minutes {
			continue
		}

		mayPush(&stockpile_res)

		r_geode := blueprint.Robots[utilities.ResGeode]
		if r_geode.CanProduce(&state.Items) {
			buildRobot(state, r_geode.Output)
			continue
		}

		for i := 0; i < utilities.TotalResources; i++ {
			robot := blueprint.Robots[i]
			if robot.CanProduce(&state.Items) {
				buildRobot(state, robot.Output)
				continue
			}
		}
	}

	fmt.Printf("Blueprint %d yields %d\n", blueprint.ID, max_geodes)

	return max_geodes
}

func PartB(blueprints []utilities.Blueprint) int {
	const max_minutes = 32
	product := 1

	num_blueprints := utilities.MinInt(3, len(blueprints))

	for idx := 0; idx < num_blueprints; idx++ {
		bp := blueprints[idx]
		product *= MaxGeodes(bp, max_minutes)
	}

	return product
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
