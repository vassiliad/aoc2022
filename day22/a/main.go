package main

import (
	"fmt"
	"os"

	"github.com/vassiliad/aoc2022/day22/a/utilities"
)

func GetAgentDeltaMovement(agent_dir utilities.Direction) (int, int) {
	dx, dy := 0, 0

	switch agent_dir {
	case utilities.DirUp:
		dy = -1
	case utilities.DirDown:
		dy = 1
	case utilities.DirLeft:
		dx = -1
	case utilities.DirRight:
		dx = 1
	}
	return dx, dy
}

func CalcAgentFace(agent_dir utilities.Direction) string {
	agent_facing := ">"

	switch agent_dir {
	case utilities.DirUp:
		agent_facing = "^"
	case utilities.DirDown:
		agent_facing = "v"
	case utilities.DirLeft:
		agent_facing = "<"
	}

	return agent_facing
}

func PartA(chamber *utilities.Chamber) int {
	agent_pos := utilities.Point{}
	agent_dir := utilities.DirRight

	// VV: The agent got the shmoves
	shmoves := map[utilities.Point]string{}

	// VV: First, pinpoint starting point. It's the left-most `Floor` tile in height = 0
	// (Highest points in map have height = 0, lowest points have height = chamber.Height-1)

	for x := 0; x < chamber.Width; x++ {
		agent_pos.X = x
		if chamber.Tiles[agent_pos] == utilities.TileOpen {
			break
		}
	}

	shmoves[agent_pos] = CalcAgentFace(agent_dir)
	Draw(chamber, shmoves)

	for _, instr := range chamber.Instructions {
		if instr.Turn != 0 {
			agent_dir = (agent_dir + 4 + utilities.Direction(instr.Turn)) % 4
			shmoves[agent_pos] = CalcAgentFace(agent_dir)
			continue
		}

		if instr.Walk == 0 {
			continue
		}

		dx, dy := GetAgentDeltaMovement(agent_dir)

		for i := 0; i < int(instr.Walk); i++ {

			try := agent_pos
			try.X += dx
			try.Y += dy
			forward_tile := chamber.Tiles[try]
			if forward_tile == utilities.TileVoid {
				// VV: Keep going till you find a Wall or an Open tile
				wrap := try
				for {
					wrap.X = (chamber.Width + wrap.X + dx) % chamber.Width
					wrap.Y = (chamber.Height + wrap.Y + dy) % chamber.Height

					forward_tile = chamber.Tiles[wrap]
					if forward_tile == utilities.TileWall {
						break
					} else if forward_tile == utilities.TileOpen {
						try = wrap
						break
					}
				}
			}

			if forward_tile == utilities.TileWall {
				break
			}

			agent_pos = try
			shmoves[agent_pos] = CalcAgentFace(agent_dir)
		}
	}

	agent_face := CalcAgentFace(agent_dir)

	// VV: we knew the agent had the shmoves. Agent, show the world your shmoves
	Draw(chamber, shmoves)
	fmt.Printf("Agent %+v facing %s\n", agent_pos, agent_face)

	return 1000*(agent_pos.Y+1) + 4*(agent_pos.X+1) + int(agent_dir)
}

func Draw(chamber *utilities.Chamber, shmoves map[utilities.Point]string) {
	for y := 0; y < chamber.Height; y++ {
		for x := 0; x < chamber.Width; x++ {
			point := utilities.Point{X: x, Y: y}

			block := " "
			if face, ok := shmoves[point]; ok {
				block = face
			} else {
				tile := chamber.Tiles[point]
				if tile == utilities.TileWall {
					block = "#"
				} else if tile == utilities.TileOpen {
					block = "."
				}
			}
			fmt.Print(block)
		}
		fmt.Println()
	}

	fmt.Println("---")
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
