package main

import (
	"container/heap"
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day16/a/utilities"
)

type DontCare int8

type DjikstraPoint struct {
	ValveUID      int
	MinutesSpent  int
	Visited       utilities.BitMask
	OpenValves    utilities.BitMask
	MinimumLiquid int
}

func (d *DjikstraPoint) Report(network *utilities.Network) string {
	valve_name := network.BookReverse[d.ValveUID]

	visited_names := []string{}
	open_names := []string{}

	for i := 0; i < len(network.Valves); i++ {
		if d.Visited.Get(i) {
			visited_names = append(visited_names, network.BookReverse[i])
		}

		if d.OpenValves.Get(i) {
			open_names = append(open_names, network.BookReverse[i])
		}
	}

	return fmt.Sprintf("{Valve: %s, MinLiquid: %d, Mins: %d, Visited: %v, OpenValves: %v}",
		valve_name, d.MinimumLiquid, d.MinutesSpent, visited_names, open_names)
}

func Neighbours(network *utilities.Network, state *DjikstraPoint, max_minutes int) []DjikstraPoint {
	ret := []DjikstraPoint{}
	if state.MinutesSpent == max_minutes {
		return ret
	}

	valve_name, ok := network.BookReverse[state.ValveUID]
	if !ok {
		panic(state)
	}

	valve, ok := network.Valves[valve_name]

	if !ok {
		panic(state)
	}

	neighbours := valve.GetNeighbourUIDS(network)

	for _, u := range neighbours {
		point := *state
		point.ValveUID = u
		ret = append(ret, point)
	}

	return ret
}

type FilteredState struct {
	ValveUID int
	// MinutesSpent int
	OpenValves utilities.BitMask

	// FlowRate     int
	// Visited utilities.BitMask
}

func FilterState(point *DjikstraPoint) FilteredState {
	return FilteredState{
		OpenValves: point.OpenValves,
		ValveUID:   point.ValveUID,
		// MinutesSpent: point.MinutesSpent,
		// Visited:      point.Visited,
	}
}

type Book map[FilteredState]*utilities.Item

func (b *Book) Update(item *utilities.Item) {
	point := item.Value.(DjikstraPoint)

	(*b)[FilterState(&point)] = item
}

func Djikstra(network *utilities.Network, max_minutes int) int {
	const start_valve = "AA"

	var all_valves uint64 = 1<<uint64(len(network.Valves)) - 1

	svalve, ok := network.Valves[start_valve]

	if !ok {
		panic(network)
	}

	queue := make(utilities.PriorityQueue, 0)
	book := Book{}

	start := DjikstraPoint{
		OpenValves:    0,
		Visited:       0,
		ValveUID:      svalve.UID,
		MinimumLiquid: 0,
		MinutesSpent:  0,
	}

	item := utilities.Item{Value: start, Priority: 0}

	heap.Push(&queue, &item)
	book.Update(&item)

	max_liquid := 0

	step := uint64(0)

	for queue.Len() > 0 {
		pop := heap.Pop(&queue).(*utilities.Item)
		point := pop.Value.(DjikstraPoint)
		cvalve := network.Valves[network.BookReverse[point.ValveUID]]

		max_liquid = utilities.MaxInt(max_liquid, point.MinimumLiquid)

		if point.MinutesSpent == max_minutes || point.OpenValves == utilities.BitMask(all_valves) {
			continue
		}

		point.MinutesSpent++

		// point.Visited.Set(point.ValveUID)
		// if step%1000 == 0 {
		// 	fmt.Printf("Visiting %s maxLiquid: %d queue: %d\n",
		// 		point.Report(network), max_liquid, len(queue))
		// }
		step++

		if !point.OpenValves.Get(point.ValveUID) && cvalve.Rate > 0 {
			state_open := point

			state_open.OpenValves.Set(point.ValveUID)
			state_open.MinimumLiquid += cvalve.Rate * (max_minutes - state_open.MinutesSpent)
			max_liquid = utilities.MaxInt(max_liquid, point.MinimumLiquid)

			next := utilities.Item{
				Value:    state_open,
				Priority: -state_open.MinimumLiquid,
			}
			heap.Push(&queue, &next)
			book.Update(&next)
		}

		neighbours := Neighbours(network, &point, max_minutes)

		for _, ns := range neighbours {
			fs := FilterState(&ns)

			if old_item, ok := book[fs]; ok {
				old := old_item.Value.(DjikstraPoint)

				if (old.MinimumLiquid >= ns.MinimumLiquid) && (old.MinutesSpent <= ns.MinutesSpent) {
					continue
				}
			}

			ni := utilities.Item{Value: ns, Priority: -ns.MinimumLiquid}
			heap.Push(&queue, &ni)
			book.Update(&ni)
		}
	}

	return max_liquid
}

func PartA(pairs *utilities.Network) int {
	return Djikstra(pairs, 30)
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
