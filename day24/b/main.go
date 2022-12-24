package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/vassiliad/aoc2022/day24/a/utilities"
)

// VV: Every LeastCommonMultiple(width, Height) the blizzards reset to their original positions.
// Therefore the "hash" State of the agent is position of the "Agent" and the minutes since the last
// time the Blizzards where in their original positions
type Memo struct {
	PeriodTime int
	Pos        int
}

type State struct {
	Time    int
	EstTime int
	Memo
}

const (
	ToRight = 0
	ToLeft  = 1
	ToDown  = 2
	ToUp    = 3
)

type Valley struct {
	Visited map[Memo]int8
	Queue   utilities.PriorityQueue
	Prev    map[State]State
	// VV: PeriodTime, y, Bitmask uses 4 bits to represent the 4 wind-directions
	Blizards      map[int]map[int]map[int]utilities.BitMask
	LCMDimensions int
	utilities.Valley
}

// VV: I thought my code that was messing about with indices had a bug so I wrote this alternative approach.
// Turns out the bug was not in that code. I just forgot to check that "staying" in the same place requires
// that winds do not move in on the same spot as the Elf. Those warm clothes are no good ...
func (v *Valley) InitBlizards() {
	v.LCMDimensions = utilities.ComputeLeastCommonMultiple([]int{v.Width - 2, v.Height})

	v.Blizards = map[int]map[int]map[int]utilities.BitMask{}

	// VV: Initial state
	blizard := map[int]map[int]utilities.BitMask{}
	for y := 0; y < v.Height; y++ {
		blizard[y] = map[int]utilities.BitMask{}
		for x := 1; x < v.Width-1; x++ {
			b := utilities.BitMask(0)
			if v.BlizRight[y][x] == 1 {
				b.Set(ToRight)
			}

			if v.BlizLeft[y][x] == 1 {
				b.Set(ToLeft)
			}

			if v.BlizDown[y][x] == 1 {
				b.Set(ToDown)
			}

			if v.BlizUp[y][x] == 1 {
				b.Set(ToUp)
			}

			blizard[y][x] = b
		}
	}
	v.Blizards[0] = blizard

	for p := 0; p < v.LCMDimensions; p++ {
		last := v.Blizards[p]
		blizard := map[int]map[int]utilities.BitMask{}
		for y := 0; y < v.Height; y++ {
			blizard[y] = map[int]utilities.BitMask{}
			for x := 1; x < v.Width; x++ {
				b := utilities.BitMask(0)

				l_right := last[y][x+1]
				l_left := last[y][x-1]
				l_down := last[y+1][x]
				l_up := last[y-1][x]

				if x == 1 {
					l_left |= last[y][v.Width-2]
				}

				if x == v.Width-2 {
					l_right |= last[y][1]
				}

				if y == 0 {
					l_up |= last[v.Height-1][x]
				}

				if y == v.Height-1 {
					l_down |= last[0][x]
				}

				if l_right.Get(ToLeft) {
					b.Set(ToLeft)
				}

				if l_left.Get(ToRight) {
					b.Set(ToRight)
				}

				if l_up.Get(ToDown) {
					b.Set(ToDown)
				}

				if l_down.Get(ToUp) {
					b.Set(ToUp)
				}

				blizard[y][x] = b
			}
		}
		v.Blizards[p+1] = blizard
	}
}

func (s *State) EstimateDistanceToTargetPosition(v *Valley, target int) {
	x := (s.Pos + v.Width) % v.Width
	y := s.Pos / v.Width

	if s.Pos < 0 {
		y = -1
	}

	t_x := (target + v.Width) % v.Width
	t_y := target / v.Width

	if target < 0 {
		t_y = -1
	}

	s.EstTime = utilities.AbsInt(t_x-x) + utilities.AbsInt(t_y-y) + s.Time
}

func (v *Valley) IndexLeftX(x, time int) int {
	return (x-1+v.Width-2+time)%(v.Width-2) + 1
}

func (v *Valley) IndexRightX(x, time int) int {
	return (x-1+v.Width-2-(time%(v.Width-2)))%(v.Width-2) + 1
}

func (v *Valley) IndexUpY(y, time int) int {
	return (y + time) % v.Height
}

func (v *Valley) IndexDownY(y, time int) int {
	return (y + v.Height - (time % v.Height)) % v.Height
}

func (v *Valley) Draw(position, time int, print_too bool) string {

	px := (position + v.Width) % v.Width
	py := position / v.Width

	if position < 0 {
		py = -1
	}

	ret := fmt.Sprintf("Round %d: %d=(%d, %d)\n", time, position, px, py)

	// return ret

	for x := 0; x < v.Width; x++ {
		if px == x && py == -1 {
			ret += "E"
		} else if x == v.Start {
			ret += "."
		} else {
			ret += "#"
		}
	}

	ret += "\n"

	toInt := func(b bool) int8 {
		if b {
			return 1
		}

		return 0
	}

	for y := 0; y < v.Height; y++ {
		ret += "#"

		for x := 1; x < v.Width-1; x++ {
			if px == x && py == y {
				ret += "E"
			} else {
				b := v.Blizards[time%v.LCMDimensions][y][x]

				multiple := toInt(b.Get(ToRight)) + toInt(b.Get(ToLeft)) + toInt(b.Get(ToUp)) + toInt(b.Get(ToDown))

				if multiple > 1 {
					ret += fmt.Sprint(multiple)
				} else if b.Get(ToRight) {
					ret += ">"
				} else if b.Get(ToLeft) {
					ret += "<"
				} else if b.Get(ToDown) {
					ret += "v"
				} else if b.Get(ToUp) {
					ret += "^"
				} else {
					ret += "."
				}
			}
		}
		ret += "#\n"
	}

	for x := 0; x < v.Width; x++ {
		if px == x && py == v.Height {
			ret += "E"
		} else if x == v.End {
			ret += "."
		} else {
			ret += "#"
		}
	}

	ret += "\n\n"

	if print_too {
		fmt.Print(ret)
	}

	return ret
}

func (v *Valley) Neighbours(state *State) []int {
	ret := []int{}

	deltas := [4]int{state.Pos - v.Width, state.Pos - 1, state.Pos + 1, state.Pos + v.Width}

	t := (state.Time + 1) % v.LCMDimensions

	for _, neighbour := range deltas {
		n_x := (neighbour + v.Width) % v.Width
		n_y := neighbour / v.Width
		if neighbour < 0 {
			n_y = -1
		}

		if (n_y < 0 && n_x != v.Start) || (n_y == v.Height && n_x != v.End) ||
			n_x < 1 || n_x >= v.Width-1 || n_y > v.Height {
			continue
		}

		if v.Blizards[t][n_y][n_x] > 0 {
			continue
		}

		ret = append(ret, neighbour)
	}

	return ret
}

func (v *Valley) MayPushState(state State, prev *State, target int) bool {
	state.EstimateDistanceToTargetPosition(v, target)

	state.PeriodTime = state.Time

	if _, ok := v.Visited[state.Memo]; ok {
		return false
	}

	item := utilities.Item{
		Value:    state,
		Priority: state.EstTime,
	}

	heap.Push(&v.Queue, &item)

	if prev != nil {
		p := *prev
		v.Prev[state] = p
	}
	return true
}

func (v *Valley) PrettyNeighbours(neighbours []int) []string {
	ret := []string{}

	for _, n := range neighbours {
		x := n % v.Width
		y := n / v.Width

		ret = append(ret, fmt.Sprintf("(%d,%d)", x, y))
	}

	return ret
}

func (v *Valley) TimeToMoveTo(target int, initial State) State {
	steps := 0
	for k := range v.Visited {
		delete(v.Visited, k)
	}

	v.Queue = utilities.PriorityQueue{}

	v.MayPushState(initial, nil, target)

	for v.Queue.Len() > 0 {
		item := heap.Pop(&v.Queue).(*utilities.Item)
		state := item.Value.(State)

		future := state

		if state.Pos == target {
			return state
		}

		if _, ok := v.Visited[state.Memo]; ok {
			continue
		}

		v.Visited[state.Memo] = 1

		steps++
		neighbours := v.Neighbours(&state)

		future.Time++

		n_x := (state.Pos + v.Width) % v.Width
		n_y := state.Pos / v.Width

		if state.Pos < 0 {
			n_y = -1
		}

		if v.Blizards[future.Time%v.LCMDimensions][n_y][n_x] == 0 {
			v.MayPushState(future, &state, target)
		}

		for _, n := range neighbours {
			next := future
			next.Pos = n
			next.EstimateDistanceToTargetPosition(v, target)
			v.MayPushState(next, &state, target)
		}

	}

	return State{Time: -1}
}

func PartB(v *utilities.Valley) int {
	valley := Valley{
		Valley: *v, Queue: utilities.PriorityQueue{},
		Visited: map[Memo]int8{},
		Prev:    map[State]State{},
	}

	valley.InitBlizards()
	target_start := v.Start - v.Width
	target_end := v.Height*v.Width + v.End

	initial := State{
		Time: 0,
		Memo: Memo{
			Pos: target_start,
		},
	}

	initial.EstimateDistanceToTargetPosition(&valley, target_end)

	initial = valley.TimeToMoveTo(target_end, initial)
	initial = valley.TimeToMoveTo(target_start, initial)
	initial = valley.TimeToMoveTo(target_end, initial)

	return initial.Time
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
