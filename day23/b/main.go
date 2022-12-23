package main

import (
	"fmt"
	"math"
	"os"

	"github.com/vassiliad/aoc2022/day23/a/utilities"
)

type Work struct {
	Round int
	// VV: keep track of round numbers that Elf[i] stops moving
	Stopped      [][]int
	OrderedElves []utilities.Point
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
	for y := w.Top; y >= w.Bottom; y-- {
		for x := w.Left; x <= w.Right; x++ {
			p := utilities.Point{X: x, Y: y}
			if _, ok := w.Elves[p]; ok {
				w.OrderedElves = append(w.OrderedElves, p)
				w.Stopped = append(w.Stopped, []int{})
			}
		}
	}

	fmt.Println("Decided order of", len(w.OrderedElves), "elves")
}

func (w *Work) ElfDecision(elf *utilities.Point) *utilities.Point {
	// VV: An Elf checks their N, S, W, and E. Each round they change the order that they test directions
	// Moving the 1st direction to last.
	move := utilities.Direction(w.Round % 4)

	spotIsAvailableGroup := func(point *utilities.Point) bool {
		_, ok := w.Elves[*point]
		return !ok
	}

	neighbours := 0

	for y := elf.Y - 1; y <= elf.Y+1; y++ {
		for x := elf.X - 1; x <= elf.X+1; x++ {
			if _, ok := w.Elves[utilities.Point{X: x, Y: y}]; ok {
				neighbours++
			}
		}
	}

	// VV: There must be at least 1 more Elf nearby (other that myself)
	if neighbours < 2 {
		return nil
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

func (w *Work) ActRound() {
	decisions := map[utilities.Point]int{}

	for idx, elf := range w.OrderedElves {
		d := w.ElfDecision(&elf)

		if d != nil {
			if _, conflict := decisions[*d]; !conflict {
				decisions[*d] = idx
			} else {
				decisions[*d] = -1
			}
		} else {
			w.Stopped[idx] = append(w.Stopped[idx], w.Round)
		}
	}

	// VV: We need to do this in 2 steps. First remove all the old elves that made a move
	// Then add the elves to the spots that they moved into
	for _, elf_idx := range decisions {
		// VV: -1 indicates that multiple Elves decided to move to the same spot
		if elf_idx != -1 {
			delete(w.Elves, w.OrderedElves[elf_idx])
		}
	}
	for new_spot, elf_idx := range decisions {
		if elf_idx != -1 {
			if _, ok := w.Elves[new_spot]; ok {
				panic(fmt.Sprintf("Elf %+v is already there", new_spot))
			}

			w.OrderedElves[elf_idx].X, w.OrderedElves[elf_idx].Y = new_spot.X, new_spot.Y
			w.Elves[w.OrderedElves[elf_idx]] = 1
		}
	}

	w.Left, w.Right = math.MaxInt, math.MinInt
	w.Top, w.Bottom = math.MinInt, math.MaxInt

	for _, elf := range w.OrderedElves {
		w.Left = utilities.MinInt(w.Left, elf.X)
		w.Bottom = utilities.MinInt(w.Bottom, elf.Y)

		w.Right = utilities.MaxInt(w.Right, elf.X)
		w.Top = utilities.MaxInt(w.Top, elf.Y)
	}

	w.Round++
}

func PartB(group *utilities.Group) int {
	const rounds = 10

	w := Work{
		Group:        *group,
		Stopped:      [][]int{},
		OrderedElves: []utilities.Point{},
	}

	w.Init()

	w.Draw(true)

	for i := 1; i <= rounds; i++ {
		w.ActRound()
	}

	w.Draw(true)

	fmt.Println("Top", w.Top, "bottom", w.Bottom, "Left", w.Left, "Right", w.Right, "Elves", len(w.Elves))
	fmt.Printf("%+v\n", w.Stopped)

	return 0
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
