package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int8
type Tile int8

const (
	DirRight Direction = 0
	DirDown  Direction = 1
	DirLeft  Direction = 2
	DirUp    Direction = 3
	TileVoid Tile      = 0
	TileOpen Tile      = 1
	TileWall Tile      = 2
)

type Instruction struct {
	// VV: Direction goes Up, Right, Down, Left.
	// So Turn = -1 turns left, and Turn = +1 turns right
	// Turn = 0 means walk forward
	Turn int8
	Walk uint8
}

type Point struct {
	X int
	Y int
}

type Chamber struct {
	Width  int
	Height int

	Tiles        map[Point]Tile
	Instructions []Instruction
}

func (c *Chamber) InstructionsToString() string {
	ret := ""
	for _, instr := range c.Instructions {
		if instr.Turn == 0 {
			ret += fmt.Sprintf("%d", instr.Walk)
		} else if instr.Turn == -1 {
			ret += "L"
		} else {
			ret += "R"
		}
	}

	return ret
}

func ReadScanner(scanner *bufio.Scanner) (*Chamber, error) {
	ret := Chamber{Tiles: map[Point]Tile{}}

	parsing_instructions := false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, " \n")

		if len(line) == 0 {
			if ret.Height > 0 {
				parsing_instructions = true
			}
			continue
		}

		if !parsing_instructions {
			begin := strings.LastIndex(line, " ")

			end := len(line)

			for x := begin + 1; x < end; x++ {
				point := Point{X: x, Y: ret.Height}

				tile := TileOpen
				if line[x] == '#' {
					tile = TileWall
				} else if line[x] != '.' {
					return &ret, fmt.Errorf("unexpected '%c' at index %d in line %d \"%s\"", line[x], x, ret.Height, line)
				}
				ret.Tiles[point] = tile
			}

			ret.Width = MaxInt(ret.Width, end)
			ret.Height++
		} else {
			moved := 0
			for len(line) > 0 {
				next := len(line)

				next_turn_right := strings.Index(line, "R")
				next_turn_left := strings.Index(line, "L")

				if next_turn_right != -1 && (next_turn_left == -1 || next_turn_right < next_turn_left) {
					next = next_turn_right
				}

				if next_turn_left != -1 && (next_turn_right == -1 || next_turn_left < next_turn_right) {
					next = next_turn_left
				}

				if next == next_turn_left && next == 0 {
					ret.Instructions = append(ret.Instructions, Instruction{Turn: -1})
					next++
				} else if next == next_turn_right && next == 0 {
					ret.Instructions = append(ret.Instructions, Instruction{Turn: 1})
					next++
				} else {
					steps, err := strconv.Atoi(line[:next])

					if err != nil {
						return &ret, fmt.Errorf("unable to decode steps in \"%s\" (chars %d)", line, next)
					}

					ret.Instructions = append(ret.Instructions, Instruction{Walk: uint8(steps)})
				}

				moved += next
				line = line[next:]
			}

		}

	}

	return &ret, scanner.Err()
}

func ReadString(text string) (*Chamber, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Chamber, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
