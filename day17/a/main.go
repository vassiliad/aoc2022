package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day17/a/utilities"
)

type Tile int8

const (
	TileEmpty Tile = 0
	TileRock  Tile = 1
)

type Sky struct {
	Space  map[int]Tile
	Height int
	Width  int
}

func NewSky() Sky {
	return Sky{Space: map[int]Tile{}, Height: 0, Width: 9}
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

func (s *Sky) SimulateRockFall(index int, rock utilities.Rock, jets []utilities.Jet, jet_idx int) int {
	stencil := rock.GetColliderIndex(index, s.Width)
	for ; true; jet_idx = (jet_idx + 1) % len(jets) {
		dx := -1
		if jets[jet_idx] == utilities.JetRight {
			dx = 1
		}

		stencil, _ = s.CanMoveRock(stencil, dx, 0)
		if st, valid := s.CanMoveRock(stencil, 0, -1); valid {
			stencil = st
		} else {
			jet_idx++
			break
		}
	}

	s.RecordRockIndex(stencil)

	return jet_idx
}

func PartA(jets []utilities.Jet) int {
	sky := NewSky()

	jet_idx := 0

	for rocks := 0; rocks < 2022; rocks++ {
		index := (sky.Height+4)*sky.Width + 2 + 1
		jet_idx = sky.SimulateRockFall(index, utilities.Rock(rocks%utilities.RockNumber), jets, jet_idx)
	}

	return sky.Height
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
