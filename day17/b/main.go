package main

import (
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day17/a/utilities"
)

type Tile int8

const (
	TileEmpty Tile = 0
	TileRock  Tile = 1
)

type State struct {
	JetIndex  int
	RockIndex int
}

type SkyBase struct {
	Width                    int
	Height                   int
	CurrIterRock             utilities.Rock
	TotalRocks               int
	TimesCycleRepeated       int
	RepeatsEveryRocks        int
	RepeatsStartingFromRocks int
	RepeatHeight             int
	HeightBefore1stCycle     int
}

type Sky struct {
	Space             map[int]Tile
	RockIndexToHeight map[int]int
	HeightToRockIndex map[int]int
	TotalJets         int
	SkyBase
	State
}

type SkyBook struct {
	Book map[State]Sky
	Sky
}

func NewSky() SkyBook {
	sky := SkyBook{
		Sky: Sky{
			Space:             map[int]Tile{},
			HeightToRockIndex: map[int]int{},
			RockIndexToHeight: map[int]int{},
		},
		Book: map[State]Sky{},
	}

	sky.CurrIterRock = utilities.RockA
	sky.Width = 7 + 2

	return sky
}

func (s *Sky) IsOccupiedIndex(index int) bool {
	x, y := index%s.Width, index/s.Width

	if x == 0 || x == s.Width-1 {
		return true
	}

	if y < 1 || x < 0 {
		return true
	}

	other, ok := s.Space[index]
	return ok && other == TileRock
}

func (s *Sky) RecordRockIndex(stencil []int) {
	for _, delta := range stencil {
		y := delta / s.Width
		s.Height = utilities.MaxInt(s.Height, y)

		// s.RockIndexToHeight[s.CurrNumRock-1] = y
		s.HeightToRockIndex[y] = s.TotalRocks - 1

		s.Space[delta] = TileRock
	}
}

func (s *Sky) CanMoveRock(stencil []int, dx, dy int) ([]int, bool) {
	for i := 0; i < len(stencil); i++ {
		if s.IsOccupiedIndex(stencil[i] + dx + dy*s.Width) {
			return stencil, false
		}
	}

	for i := 0; i < len(stencil); i++ {
		stencil[i] += dx + dy*s.Width
	}

	return stencil, true
}

func (s *SkyBook) VerifyCycle(repeat_cycle int) bool {
	for dy := 0; dy <= repeat_cycle; dy++ {
		top_y := s.Height - dy
		bottom_y := s.Height - dy - repeat_cycle

		for x := 1; x < s.Width-1; x++ {
			if s.Space[x+top_y*s.Width] != s.Space[x+bottom_y*s.Width] {
				return false
			}
		}
	}

	return true
}

func (s *SkyBook) SimulateRockFall(jets []utilities.Jet, detect_loops bool) bool {
	rock := s.CurrIterRock
	spawn_index := (s.Height+4)*s.Width + 2 + 1

	stencil := rock.GetColliderIndex(spawn_index, s.Width)
	for {
		dx := -1
		if jets[s.JetIndex] == utilities.JetRight {
			dx = 1
		}
		stencil, _ = s.CanMoveRock(stencil, dx, 0)
		s.TotalJets++
		s.JetIndex = (s.JetIndex + 1) % len(jets)

		if st, valid := s.CanMoveRock(stencil, 0, -1); valid {
			stencil = st
		} else {
			break
		}
	}

	s.RecordRockIndex(stencil)
	s.CurrIterRock = (s.CurrIterRock + 1) % utilities.Rock(utilities.RockNumber)
	s.RockIndex = (s.RockIndex + 1) % utilities.RockNumber
	s.RockIndexToHeight[s.TotalRocks] = s.Height
	s.TotalRocks++

	SetupNextRock := func() {
		s.Book[s.State] = s.Sky
	}

	defer SetupNextRock()

	if detect_loops && s.TotalRocks > 2*utilities.RockNumber && s.TotalJets >= len(jets) {
		if old, ok := s.Book[s.State]; ok {
			cycle_rocks := s.TotalRocks - old.TotalRocks
			cycle_height := s.Height - old.Height

			if s.VerifyCycle(cycle_height) {
				if s.TimesCycleRepeated == 0 {
					// VV: This looks like a repeating cycle.
					// Record the Height at the start of the end of the previous "iteration" of this cycle
					// We're shooting for PreCycleHeight + (numFullCycles)*cycleHeight + Height(remainingBlocks)
					s.HeightBefore1stCycle = s.RockIndexToHeight[old.TotalRocks-1] - cycle_height
					s.RepeatsStartingFromRocks = old.TotalRocks
					s.RepeatHeight = cycle_height
					s.RepeatsEveryRocks = cycle_rocks
				}

				s.TimesCycleRepeated++

				if s.RepeatHeight != cycle_height {
					if s.TimesCycleRepeated > 0 {
						s.TimesCycleRepeated = 0
						return false
					}
					s.TimesCycleRepeated = 0
				}

				if s.TimesCycleRepeated == 3 {
					fmt.Printf("Cycle %+v=%d every %d] %+v -> %+v\n",
						s.State, cycle_height, cycle_rocks,
						old.SkyBase, s.SkyBase)
					return true
				}
			} else {
				s.TimesCycleRepeated = 0
			}
		} else {
			s.TimesCycleRepeated = 0
		}
	}

	return false
}

func PartB(jets []utilities.Jet) int {
	const ungodly = 1000000000000

	sky := NewSky()

	was_loopy := false
	for rocks := 0; !was_loopy; rocks++ {
		was_loopy = sky.SimulateRockFall(jets, true)
	}

	num_loops := (ungodly-sky.RepeatsStartingFromRocks)/sky.RepeatsEveryRocks + 1
	remainder := (ungodly - sky.RepeatsStartingFromRocks) % sky.RepeatsEveryRocks

	fmt.Println(
		"CycleEnd", sky.TotalRocks,
		"HeightBeforeCycle", sky.HeightBefore1stCycle,
		"RepeatRocks", sky.RepeatsEveryRocks,
		"CycleHeight", sky.RepeatHeight,
		"RemainderRocks", remainder,
		"Repeated", num_loops,
	)

	all_cycles_height := num_loops * sky.RepeatHeight

	curr_height := sky.Height
	for rocks := 0; rocks < remainder; rocks++ {
		sky.SimulateRockFall(jets, false)
	}
	remainder_height := sky.Height - curr_height

	// VV: Formula is: PreCycleHeight + (numFullCycles)*cycleHeight + Height(remainingBlocks)
	return sky.HeightBefore1stCycle + all_cycles_height + remainder_height
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
