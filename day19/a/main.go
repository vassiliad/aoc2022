package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/vassiliad/aoc2022/day19/a/utilities"
)

type Memo struct {
	Inventory utilities.ManyParts
	Robots    utilities.ManyParts
}

type State struct {
	Minute int
	// BlueprintID int
	Memo
}

type Explore struct {
	Queue      utilities.PriorityQueue
	Visited    map[State]int
	MaxMinutes int
}

// func (e *Explore) MinimumGeodes(state *State) int {
// 	rem_time := e.MaxMinutes - state.Minute

// }

func (e *Explore) ShouldExplore(state *State) bool {
	if _, visited := e.Visited[*state]; visited {
		return false
	}

	return true
}

func (e *Explore) MayPush(state State) bool {
	if !e.ShouldExplore(&state) {
		// fmt.Printf("NO  %+v\n", state)
		return false
	}

	item := utilities.Item{
		Value: state,
		// Priority: -state.Minute,
		Priority: -int(state.Inventory[utilities.ResGeode]),
		// state.Robots[utilities.ResGeode]*1000 -
	}

	heap.Push(&e.Queue, &item)
	return true
}

func MaxGeodes(blueprint utilities.Blueprint, max_minutes int) int {
	explore := Explore{
		Queue:      make(utilities.PriorityQueue, 0),
		Visited:    make(map[State]int),
		MaxMinutes: max_minutes,
	}

	// VV: skip a round because there's no robot that requires 0 components
	init := State{Minute: 1}
	init.Robots[utilities.ResOre] = 1
	init.Inventory[utilities.ResOre] = 1

	explore.MayPush(init)

	max_geodes := 0

	// VV: We can prune the tree by avoiding to build a robot after we have "enough" of them
	max_required := utilities.ManyParts{}

	for _, r := range blueprint.Robots {
		// VV: "enough" is when you're already producing enough each minute to build
		// at least 1 of any other robot (we can only build 1 robot at a time)
		// max_required[utilities.ResOre] += r.CostOre
		// max_required[utilities.ResClay] += r.CostClay
		// max_required[utilities.ResObsidian] += r.CostObsidian

		// VV: "enough" is when you have just enough robots to satisfy the highest demand for any robot per minute
		// Since we're building just 1 robot at a time, this should be "safe"
		max_required[utilities.ResOre] = utilities.MaxInt(r.CostOre, int(max_required[utilities.ResOre]))
		max_required[utilities.ResClay] = utilities.MaxInt(r.CostClay, int(max_required[utilities.ResClay]))
		max_required[utilities.ResObsidian] = utilities.MaxInt(r.CostObsidian, int(max_required[utilities.ResObsidian]))
	}

	max_required[utilities.ResGeode] = math.MaxInt

	step := 0
	for explore.Queue.Len() > 0 {
		pop := explore.Queue.Pop().(*utilities.Item)
		state := pop.Value.(State)

		if max_geodes < int(state.Inventory[utilities.ResGeode]) {
			max_geodes = int(state.Inventory[utilities.ResGeode])
			fmt.Printf("%d -> %d] %+v\n", explore.Queue.Len(), max_geodes, state)
		}

		if step%1000000 == 0 {
			fmt.Printf("%d] %d,%d : %d] %+v\n",
				blueprint.ID, explore.Queue.Len(), len(explore.Visited),
				max_geodes, state)
		}
		step++

		if _, visited := explore.Visited[state]; visited || state.Minute == explore.MaxMinutes {
			continue
		}

		explore.Visited[state] = 1

		state.Minute++

		next := state
		for i := 0; i < utilities.TotalResources; i++ {
			next.Inventory[i] += next.Robots[i]
		}
		explore.MayPush(next)

		for _, r := range blueprint.Robots {
			if (state.Inventory[r.Output] > max_required[r.Output] && r.Output != utilities.ResGeode) || !r.CanProduce(&state.Inventory) {
				continue
			}
			next := state
			r.Pay(&next.Inventory)
			for i := 0; i < utilities.TotalResources; i++ {
				next.Inventory[i] += next.Robots[i]
			}
			next.Robots[r.Output] += 1
			explore.MayPush(next)
		}
	}

	fmt.Printf("Blueprint %d yields %d (explored %d)\n", blueprint.ID, max_geodes, len(explore.Visited))

	return max_geodes
}

func PartA(blueprints []utilities.Blueprint) int {
	const max_minutes = 24
	sum := 0

	for _, bp := range blueprints {
		sum += bp.ID * MaxGeodes(bp, max_minutes)
	}

	return sum
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
