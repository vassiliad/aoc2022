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

func SorryCalcTestCube(
	agent_pos utilities.Point,
	agent_dir utilities.Direction,
	cube_size int,
	chamber *utilities.Chamber) (
	utilities.Point, utilities.Direction) {
	// VV: agent_pos here is ALREADY in the void
	cube_x := (agent_pos.X / cube_size) % 4
	cube_y := (agent_pos.Y / cube_size) % 4

	if cube_x == 3 && cube_y == 1 {
		if agent_dir == utilities.DirRight {
			return Teleport(cube_x, cube_y+1, 90, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 2 && cube_y == 3 {
		if agent_dir == utilities.DirDown {
			return Teleport(cube_x-2, cube_y-2, 180, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 1 && cube_y == 0 {
		if agent_dir == utilities.DirUp {
			return Teleport(cube_x+1, cube_y, 90, cube_size, agent_pos, agent_dir)
		}
	}

	panic(fmt.Sprintf("Vassili, hard-code edges %d, %d with direction %d (agent %+v)",
		cube_x, cube_y, agent_dir, agent_pos))
}

func Teleport(
	cube_x, cube_y, rotate_degrees_right, cube_size int,
	agent_pos utilities.Point,
	agent_dir utilities.Direction,
) (
	utilities.Point, utilities.Direction) {
	// VV: TurnKind = 1 means turn right, -1 means turn Left
	cube_pos := utilities.Point{X: agent_pos.X % cube_size, Y: agent_pos.Y % cube_size}

	new_dir := agent_dir
	for deg := 0; deg < rotate_degrees_right; deg += 90 {
		cube_pos.X, cube_pos.Y = cube_size-cube_pos.Y-1, cube_pos.X
		new_dir++
	}

	agent_pos.X = cube_x*cube_size + cube_pos.X
	agent_pos.Y = cube_y*cube_size + cube_pos.Y

	new_dir = new_dir % 4
	return agent_pos, new_dir
}

func VassiliCalcCube(
	agent_pos utilities.Point,
	agent_dir utilities.Direction,
	cube_size int,
	chamber *utilities.Chamber) (
	utilities.Point, utilities.Direction) {
	// VV: agent_pos here is ALREADY in the void
	cube_x := (agent_pos.X / cube_size) % 4
	cube_y := (agent_pos.Y / cube_size) % 4

	if cube_x == 0 && cube_y == 0 {
		if agent_dir == utilities.DirLeft {
			return Teleport(cube_x, cube_y+2, 180, cube_size, agent_pos, agent_dir)
		} else if agent_dir == utilities.DirDown {
			return Teleport(cube_x+2, cube_y, 0, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 3 && cube_y == 2 {
		if agent_dir == utilities.DirLeft {
			return Teleport(cube_x-2, cube_y-2, 180, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 1 && cube_y == 3 {
		if agent_dir == utilities.DirDown {
			return Teleport(cube_x-1, cube_y, 90, cube_size, agent_pos, agent_dir)
		} else if agent_dir == utilities.DirUp {
			return Teleport(cube_x-1, cube_y, 90, cube_size, agent_pos, agent_dir)
		} else if agent_dir == utilities.DirRight {
			return Teleport(cube_x, cube_y-1, 270, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 2 && cube_y == 1 {
		if agent_dir == utilities.DirRight {
			return Teleport(cube_x, cube_y-1, 270, cube_size, agent_pos, agent_dir)
		} else if agent_dir == utilities.DirDown {
			return Teleport(cube_x-1, cube_y, 90, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 0 && cube_y == 1 {
		if agent_dir == utilities.DirUp {
			return Teleport(cube_x+1, cube_y, 90, cube_size, agent_pos, agent_dir)
		} else if agent_dir == utilities.DirLeft {
			return Teleport(cube_x, cube_y+1, 270, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 3 && cube_y == 3 {
		if agent_dir == utilities.DirLeft {
			return Teleport(cube_x-2, cube_y-3, 270, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 3 && cube_y == 0 {
		if agent_dir == utilities.DirRight {
			return Teleport(cube_x-2, cube_y+2, 180, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 2 && cube_y == 2 {
		if agent_dir == utilities.DirRight {
			return Teleport(cube_x, cube_y-2, 180, cube_size, agent_pos, agent_dir)
		}
	} else if cube_x == 2 && cube_y == 3 {
		if agent_dir == utilities.DirUp {
			return Teleport(cube_x-2, cube_y, 0, cube_size, agent_pos, agent_dir)
		}
	}

	panic(fmt.Sprintf("Vassili, hard-code edges %d, %d with direction %d (agent %+v)",
		cube_x, cube_y, agent_dir, agent_pos))
}

func FindNextCubeFaceToTeleportTo(
	agent_pos utilities.Point,
	agent_dir utilities.Direction,
	cube_size int,
	chamber *utilities.Chamber) (
	utilities.Point, utilities.Direction) {

	// VV: Ideally, you want to figure out a way to generate the Graph that connects the
	// edges of the cube. I *thought* I figured out a way to do it, but it didn't work.
	// It assumed
	if cube_size == 4 {
		return SorryCalcTestCube(agent_pos, agent_dir, cube_size, chamber)
	}

	// VV: I am sorry
	return VassiliCalcCube(agent_pos, agent_dir, cube_size, chamber)
}

func PartB(chamber *utilities.Chamber) int {
	agent_pos := utilities.Point{}
	agent_dir := utilities.DirRight

	cube_size := chamber.Height / 4
	longestDimension := chamber.Height

	if chamber.Width > chamber.Height {
		longestDimension = chamber.Width
		cube_size = chamber.Width / 4
	}

	fmt.Println("Cube size is", cube_size)

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

	defer func() {
		Draw(chamber, shmoves)
	}()

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
			try.X = (try.X + dx + longestDimension) % longestDimension
			try.Y = (try.Y + dy + longestDimension) % longestDimension
			forward_tile := chamber.Tiles[try]
			if forward_tile == utilities.TileVoid {
				wrap, new_dir := FindNextCubeFaceToTeleportTo(try, agent_dir, cube_size, chamber)

				forward_tile = chamber.Tiles[wrap]
				if forward_tile == utilities.TileWall {
					break
				} else if forward_tile == utilities.TileOpen {
					try = wrap
					agent_dir = new_dir
					dx, dy = GetAgentDeltaMovement(agent_dir)
				}
			}

			if forward_tile == utilities.TileWall {
				break
			}

			agent_pos = try
			shmoves[agent_pos] = CalcAgentFace(agent_dir)
			// Draw(chamber, shmoves)
		}
	}

	agent_face := CalcAgentFace(agent_dir)

	// VV: we knew the agent had the shmoves. Agent, show the world your shmoves
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
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
