package main

import (
	"fmt"
	"math"
	"os"

	"github.com/vassiliad/aoc2022/day23/a/utilities"
)

type Work struct {
	Round        int
	OrderedElves []utilities.Point
	Neighbours   []map[int]int8
	utilities.Group
}

func (w *Work) CountEmptySpaces() int {
	width := w.Right - w.Left + 1
	height := w.Top - w.Bottom + 1

	return (width * height) - len(w.Elves)
}

func (w *Work) Draw(to_term_too bool) string {
	report := fmt.Sprintf("Round %d, (%d, %d) -> (%d, %d)\n",
		w.Round, w.Left, w.Top, w.Right, w.Bottom)

	for y := w.Top; y >= w.Bottom; y-- {
		for x := w.Left; x <= w.Right; x++ {
			p := utilities.Point{X: x, Y: y}
			if _, ok := w.Elves[p]; ok {
				report += "#"
			} else {
				report += "."
			}
		}
		report += "\n"
	}

	if to_term_too {
		fmt.Print(report)
	}

	return report
}

func (w *Work) Init() {
	fmt.Println("There are", len(w.OrderedElves), "elves")
	ordered := 0

	for y := w.Top; y >= w.Bottom; y-- {
		for x := w.Left; x <= w.Right; x++ {
			p := utilities.Point{X: x, Y: y}
			if _, ok := w.Elves[p]; ok {
				w.OrderedElves = append(w.OrderedElves, p)
				w.Neighbours = append(w.Neighbours, map[int]int8{})
				// VV: I modified PartA to do this for us, but adding the code here too for clarity
				w.Elves[p] = ordered
				ordered++
			}
		}
	}

	for idx, elf := range w.OrderedElves {
		for other_idx, other_elf := range w.OrderedElves[idx+1:] {
			if utilities.AbsInt(other_elf.X-elf.X) <= 1 && utilities.AbsInt(other_elf.Y-elf.Y) <= 1 {
				w.Neighbours[idx][other_idx+idx+1] = 0
				w.Neighbours[other_idx+idx+1][idx] = 0
			}
		}
	}
}

func (w *Work) ElfDecision(elf *utilities.Point) *utilities.Point {
	// VV: An Elf checks their N, S, W, and E. Each round they change the order that they test directions
	// Moving the 1st direction to last.
	move := utilities.Direction(w.Round % 4)

	spotIsAvailableGroup := func(point *utilities.Point) bool {
		_, ok := w.Elves[*point]
		return !ok
	}

	for i := 0; i < 4; i++ {
		var d, o1, o2 *utilities.Point

		switch move {
		case utilities.DirN:
			d = &utilities.Point{X: elf.X, Y: elf.Y + 1}
			o1, o2 = &utilities.Point{X: d.X - 1, Y: d.Y}, &utilities.Point{X: d.X + 1, Y: d.Y}
		case utilities.DirS:
			d = &utilities.Point{X: elf.X, Y: elf.Y - 1}
			o1, o2 = &utilities.Point{X: d.X - 1, Y: d.Y}, &utilities.Point{X: d.X + 1, Y: d.Y}
		case utilities.DirE:
			d = &utilities.Point{X: elf.X + 1, Y: elf.Y}
			o1, o2 = &utilities.Point{X: d.X, Y: d.Y - 1}, &utilities.Point{X: d.X, Y: d.Y + 1}
		case utilities.DirW:
			d = &utilities.Point{X: elf.X - 1, Y: elf.Y}
			o1, o2 = &utilities.Point{X: d.X, Y: d.Y - 1}, &utilities.Point{X: d.X, Y: d.Y + 1}
		}

		if spotIsAvailableGroup(o1) && spotIsAvailableGroup(o2) && spotIsAvailableGroup(d) {
			return d
		}
		move = (move + 1) % 4
	}

	return nil
}

func (w *Work) ActRound() bool {
	moved := 0
	decisions := map[utilities.Point]int{}

	for idx, elf := range w.OrderedElves {
		if len(w.Neighbours[idx]) == 0 {
			continue
		}

		d := w.ElfDecision(&elf)

		if d != nil {
			if _, conflict := decisions[*d]; !conflict {
				decisions[*d] = idx
				moved++
			} else {
				decisions[*d] = -1
			}
		}
	}

	// VV: We need to do this in 2 steps. First remove all the old elves that made a move
	// Then add the elves to the spots that they moved into
	for _, idx := range decisions {
		// VV: -1 indicates that multiple Elves decided to move to the same spot
		if idx != -1 {
			// VV: Goodbye old neighbours, I am moving out
			delete(w.Elves, w.OrderedElves[idx])
			for other_idx := range w.Neighbours[idx] {
				if _, ok := w.Neighbours[other_idx][idx]; !ok {
					panic(fmt.Sprintf("My neigbhour %d does not know me %d", other_idx, idx))
				}

				if _, ok := w.Neighbours[idx][other_idx]; !ok {
					panic(fmt.Sprintf("I (%d) do not know my neighbour %d", idx, other_idx))
				}

				delete(w.Neighbours[other_idx], idx)
				delete(w.Neighbours[idx], other_idx)
			}
		}
	}
	for new_spot, idx := range decisions {
		if idx != -1 {
			if _, ok := w.Elves[new_spot]; ok {
				panic(fmt.Sprintf("Elf %+v is already there", new_spot))
			}

			// VV: Hello neighbours, I just moved in.
			w.OrderedElves[idx].X, w.OrderedElves[idx].Y = new_spot.X, new_spot.Y

			for y := new_spot.Y - 1; y <= new_spot.Y+1; y++ {
				for x := new_spot.X - 1; x <= new_spot.X+1; x++ {
					if other_idx, ok := w.Elves[utilities.Point{X: x, Y: y}]; ok {
						w.Neighbours[idx][other_idx] = 1
						w.Neighbours[other_idx][idx] = 1
					}
				}
			}

			w.Elves[w.OrderedElves[idx]] = idx
		}
	}

	// VV: Keep track of rectangle containing Elves just to watch them move!
	w.Left, w.Right = math.MaxInt, math.MinInt
	w.Top, w.Bottom = math.MinInt, math.MaxInt

	for _, elf := range w.OrderedElves {
		w.Left = utilities.MinInt(w.Left, elf.X)
		w.Bottom = utilities.MinInt(w.Bottom, elf.Y)

		w.Right = utilities.MaxInt(w.Right, elf.X)
		w.Top = utilities.MaxInt(w.Top, elf.Y)
	}

	w.Round++

	return moved > 0
}

func PartB(group *utilities.Group) int {

	w := Work{
		Group:        *group,
		OrderedElves: []utilities.Point{},
		Neighbours:   []map[int]int8{},
	}

	w.Init()

	w.Draw(true)

	// VV: initially, I tried testing whether the Elves stop with some Period but I couldn't figure that out
	// so in the end I just trimmed down the checks I need to make by keeping track of which Elves are guaranteed
	// not to make a move during a round. These are Elves with no neighbours.
	for w.ActRound() {
	}

	w.Draw(true)

	fmt.Println("Top", w.Top, "bottom", w.Bottom, "Left", w.Left, "Right", w.Right, "Elves", len(w.Elves))

	return w.Round
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
