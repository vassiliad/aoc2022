package main

import (
	"container/heap"
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day16/a/utilities"
)

type FilteredState struct {
	ValveUID       int
	ValveUIDMinion int
	Flow           int
	MinutesSpent   int
}

type DjikstraPoint struct {
	MinimumLiquid    int
	OpenValvesMe     utilities.BitMask
	OpenValvesMinion utilities.BitMask
	Visited          utilities.BitMask
	VisitedMinion    utilities.BitMask
	Flow             int
	FilteredState
}

func (f *DjikstraPoint) IsValveOpen(uid int) bool {
	return f.OpenValvesMe.Get(uid) || f.OpenValvesMinion.Get(uid)
}

type Network struct {
	utilities.Network
}

func (n *Network) GetValve(uid int) *utilities.Valve {
	return n.Valves[n.BookReverse[uid]]
}

func (d *DjikstraPoint) Report(network *utilities.Network) string {
	valve_name := network.BookReverse[d.ValveUID]
	valve_minion_name := network.BookReverse[d.ValveUIDMinion]

	visited_names := []string{}
	open_names := []string{}

	for i := 0; i < len(network.Valves); i++ {
		if d.Visited.Get(i) || d.VisitedMinion.Get(i) {
			visited_names = append(visited_names, network.BookReverse[i])
		}

		if d.IsValveOpen(i) {
			open_names = append(open_names, network.BookReverse[i])
		}
	}

	return fmt.Sprintf("{Valve: %s/%s, Liquid: %d/%d, Mins: %d, Visited: %v, OpenValves: %v}",
		valve_name, valve_minion_name, d.MinimumLiquid, d.Flow, d.MinutesSpent, visited_names, open_names)
}

func FilterState(point *DjikstraPoint) FilteredState {
	return point.FilteredState
}

type Book map[FilteredState]*utilities.Item

func (b *Book) Update(item *utilities.Item) {
	point := item.Value.(DjikstraPoint)

	(*b)[FilterState(&point)] = item
}

func MultiDjikstra(network *Network, max_minutes int) int {
	const start_valve = "AA"

	all_valves := utilities.BitMask(0)
	num_valves := 0
	for _, v := range network.Valves {
		if v.Rate > 0 {
			all_valves.Set(v.UID)
			num_valves++
		}
	}

	fmt.Println("Maximum helpful valves", num_valves)

	svalve, ok := network.Valves[start_valve]

	if !ok {
		panic(network)
	}

	queue := make(utilities.PriorityQueue, 0)
	book := Book{}

	start := DjikstraPoint{
		FilteredState: FilteredState{
			ValveUID:       svalve.UID,
			ValveUIDMinion: svalve.UID,
		},
	}
	start.MinutesSpent = 1
	start.Visited.Set(start.ValveUID)
	start.VisitedMinion.Set(start.ValveUID)

	item := utilities.Item{Value: start, Priority: 0}

	heap.Push(&queue, &item)
	book.Update(&item)

	max_liquid := 0

	step := uint64(0)
	all_valves_open_at := -1

	IsDefinitelyWorse := func(point *DjikstraPoint, reject_equal bool) bool {
		if old_item, ok := book[point.FilteredState]; ok {
			old := old_item.Value.(DjikstraPoint)

			if (old.MinimumLiquid > point.MinimumLiquid) && (old.MinutesSpent <= point.MinutesSpent) {
				return true
			}

			if reject_equal && (old.MinimumLiquid == point.MinimumLiquid) && (old.MinutesSpent <= point.MinutesSpent) {
				return true
			}
		}
		return false
	}

	for queue.Len() > 0 {
		pop := heap.Pop(&queue).(*utilities.Item)
		point := pop.Value.(DjikstraPoint)

		max_liquid = utilities.MaxInt(max_liquid, point.MinimumLiquid)

		if IsDefinitelyWorse(&point, false) {
			continue
		}

		book.Update(pop)

		curr_valves := point.OpenValvesMe | point.OpenValvesMinion

		if curr_valves == all_valves {
			if all_valves_open_at == -1 || all_valves_open_at > point.MinutesSpent {
				all_valves_open_at = point.MinutesSpent
				fmt.Println("All valves open, current best", all_valves_open_at)
			}

			continue
		}

		if point.MinutesSpent == max_minutes-1 {
			continue
		}

		if step%100000 == 0 {
			fmt.Printf("%d-%d] %s\n", max_liquid, len(queue), point.Report(&network.Network))
		}

		step++
		valve_mine := network.GetValve(point.ValveUID)
		valve_minion := network.GetValve(point.ValveUIDMinion)

		n_mine := valve_mine.GetNeighbourUIDS(&network.Network)
		n_minion := valve_minion.GetNeighbourUIDS(&network.Network)

		doPush := func(next DjikstraPoint) {
			item := utilities.Item{Value: next, Priority: -next.MinimumLiquid}
			heap.Push(&queue, &item)
			book.Update(&item)
		}

		doTrim := func() {
			for i := 0; i < queue.Len(); i++ {
				if i != item.Index {
					c := queue[i].Value.(DjikstraPoint)

					if IsDefinitelyWorse(&c, false) {
						heap.Remove(&queue, i)
						i--
					}
				}
			}
		}

		point.MinutesSpent++

		/*VV: Neighbouring states are:
		1. Open my valve, move minion to one of its neighbours
		2. Open minion valve, I move to one of my neighbours
		3. We both open our valves (and stay here)
		3. We both move to our respective neighbours
		*/

		// VV: 1. Open my valve, move minion to one of its neighbours
		if valve_mine.Rate > 0 && !point.IsValveOpen(point.ValveUID) {
			base := point

			base.OpenValvesMe.Set(base.ValveUID)
			base.MinimumLiquid += valve_mine.Rate * (max_minutes - point.MinutesSpent)
			base.Flow += valve_mine.Rate
			if max_liquid < base.MinimumLiquid {
				max_liquid = base.MinimumLiquid
				doTrim()
			}

			for _, n := range n_minion {
				next := base

				next.ValveUIDMinion = n
				next.VisitedMinion.Set(n)

				if IsDefinitelyWorse(&next, true) {
					continue
				}

				doPush(next)
			}
		}

		// VV: 2. Open minion valve, I move to one of my neighbours
		if (valve_minion.Rate > 0) && (!point.IsValveOpen(point.ValveUIDMinion)) {
			if point.IsValveOpen(point.ValveUIDMinion) {
				panic(point)
			}
			base := point

			base.MinimumLiquid += valve_minion.Rate * (max_minutes - point.MinutesSpent)
			base.Flow += valve_minion.Rate

			base.OpenValvesMinion.Set(base.ValveUIDMinion)

			if max_liquid < base.MinimumLiquid {
				max_liquid = base.MinimumLiquid
				doTrim()
			}

			for _, n := range n_mine {
				next := base

				next.ValveUID = n
				next.Visited.Set(n)

				if IsDefinitelyWorse(&next, true) {
					continue
				}

				doPush(next)
			}
		}

		// 3. We both open our valves (and stay here)
		if valve_mine.Rate > 0 && !point.IsValveOpen(point.ValveUID) &&
			valve_minion.Rate > 0 && !point.IsValveOpen(point.ValveUIDMinion) &&
			valve_mine.UID != valve_minion.UID {
			base := point

			base.MinimumLiquid += (valve_minion.Rate + valve_mine.Rate) * (max_minutes - point.MinutesSpent)
			base.Flow += valve_minion.Rate + valve_mine.Rate

			base.OpenValvesMe.Set(base.ValveUID)
			base.OpenValvesMinion.Set(base.ValveUIDMinion)

			next := base

			if IsDefinitelyWorse(&next, true) {
				continue
			}

			doPush(next)

			if max_liquid < base.MinimumLiquid {
				max_liquid = base.MinimumLiquid
				doTrim()
			}
		}

		// VV: 4. We both move to our respective neighbours
		base := point
		for _, n := range n_mine {
			next := base
			next.ValveUID = n
			next.Visited.Set(n)

			for _, o := range n_minion {
				next := next
				next.ValveUIDMinion = o
				next.VisitedMinion.Set(o)

				if IsDefinitelyWorse(&next, true) {
					continue
				}

				doPush(next)
			}

			if len(n_minion) == 0 {
				panic(n_minion)
			}
		}

	}

	return max_liquid
}

func PartB(network *utilities.Network) int {
	my_net := Network{Network: *network}
	return MultiDjikstra(&my_net, 27)
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
