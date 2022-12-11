package main

import (
	"fmt"
	"os"

	utilities "github.com/vassiliad/aoc2022/day09/a/utilities"
)

type Board map[int]map[int]int8

func RecordTailPos(board *Board, tail_x, tail_y int) {
	if board == nil {
		return
	}

	if inner, ok := (*board)[tail_y]; ok {
		inner[tail_x] = 0
	} else {
		(*board)[tail_y] = map[int]int8{tail_x: 0}
	}
}

func DeltaStep(dstep, value int) int {
	if dstep < 0 {
		value -= 1
	} else if dstep > 0 {
		value += 1
	}

	return value
}

func SimulateRopeHead(direction utilities.Direction, x, y int) (int, int) {
	if direction == utilities.DirectionLeft {
		x--
	} else if direction == utilities.DirectionRight {
		x++
	} else if direction == utilities.DirectionUp {
		y++
	} else if direction == utilities.DirectionDown {
		y--
	} else {
		panic(direction)
	}

	return x, y
}

func SimulateRope(move *utilities.Move, board *Board, rope []int) (int, int) {
	num_knots := len(rope) / 2

	tail_x, tail_y := 0, 0

	for i := 0; i < move.Steps; i++ {
		head_x, head_y := rope[0], rope[1]
		head_x, head_y = SimulateRopeHead(move.Direction, head_x, head_y)
		rope[0], rope[1] = head_x, head_y

		for knot_idx := 0; knot_idx < (num_knots - 1); knot_idx++ {
			last_knot := knot_idx == (num_knots - 2)

			head_x, head_y := rope[knot_idx*2], rope[knot_idx*2+1]
			tail_x, tail_y = rope[knot_idx*2+2], rope[knot_idx*2+3]

			dx := head_x - tail_x
			dy := head_y - tail_y
			far_away := (utilities.AbsInt(dx) > 1 || utilities.AbsInt(dy) > 1)

			if utilities.AbsInt(dx) > 1 || far_away {
				tail_x = DeltaStep(dx, tail_x)
			}

			if utilities.AbsInt(dy) > 1 || far_away {
				tail_y = DeltaStep(dy, tail_y)
			}

			rope[knot_idx*2+2], rope[knot_idx*2+3] = tail_x, tail_y

			if last_knot {
				RecordTailPos(board, tail_x, tail_y)
			}
		}
	}

	return tail_x, tail_y
}

func RenderBoard(board *Board, min_x, max_x, min_y, max_y int) {
	for y := max_y; y >= min_y; y-- {
		if inner, ok := (*board)[y]; ok {
			for x := min_x; x <= max_x; x++ {
				if _, ok := inner[x]; ok {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func PartB(moves []utilities.Move) int {
	const num_knots = 10
	min_x, max_x, min_y, max_y := 0, 0, 0, 0

	board := Board{}

	RecordTailPos(&board, 0, 0)

	rope := make([]int, 2*num_knots)

	for _, move := range moves {
		fmt.Printf("Move: %+v\n", move)

		tail_x, tail_y := SimulateRope(&move, &board, rope)

		max_x = utilities.MaxInt(max_x, tail_x)
		max_y = utilities.MaxInt(max_y, tail_y)
		min_x = utilities.MinInt(min_x, tail_x)
		min_y = utilities.MinInt(min_y, tail_y)
	}

	println(min_x, max_x, min_y, max_y)
	RenderBoard(&board, min_x, max_x, min_y, max_y)

	positions := 0
	for _, v := range board {
		positions += len(v)
	}

	return positions
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
