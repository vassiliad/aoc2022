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

func SimulateRope(move *utilities.Move, board *Board, head_x, head_y, tail_x, tail_y int) (int, int, int, int) {
	nhead_x := head_x
	nhead_y := head_y

	for i := 0; i < move.Steps; i++ {
		if move.Direction == utilities.DirectionLeft {
			nhead_x--
		} else if move.Direction == utilities.DirectionRight {
			nhead_x++
		} else if move.Direction == utilities.DirectionUp {
			nhead_y++
		} else if move.Direction == utilities.DirectionDown {
			nhead_y--
		} else {
			panic(move)
		}

		dx := nhead_x - tail_x
		dy := nhead_y - tail_y
		far_away := (utilities.AbsInt(dx) > 1 || utilities.AbsInt(dy) > 1)

		if utilities.AbsInt(dx) > 1 || far_away {
			tail_x = DeltaStep(dx, tail_x)
		}

		if utilities.AbsInt(dy) > 1 || far_away {
			tail_y = DeltaStep(dy, tail_y)
		}

		RecordTailPos(board, tail_x, tail_y)
	}

	return nhead_x, nhead_y, tail_x, tail_y
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

func PartA(moves []utilities.Move) int {
	min_x, max_x, min_y, max_y := 0, 5, 0, 5

	head_x, head_y := 0, 0
	tail_x, tail_y := 0, 0

	board := Board{}

	RecordTailPos(&board, tail_x, tail_y)

	for _, move := range moves {
		head_x, head_y, tail_x, tail_y = SimulateRope(&move, &board, head_x, head_y, tail_x, tail_y)

		max_x = utilities.MaxInt(max_x, head_x)
		max_y = utilities.MaxInt(max_y, head_y)
		min_x = utilities.MinInt(min_x, head_x)
		min_y = utilities.MinInt(min_y, head_y)
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
	sol := PartA(input)
	logger.Println("Solution is", sol)
}
