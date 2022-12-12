package main

import (
	"container/heap"
	"math"
	"os"

	utilities "github.com/vassiliad/aoc2022/day12/a/utilities"
)

type DjikstraPoint struct {
	Previous *utilities.Item
	Position utilities.Point
}

func Neighbours(world *utilities.World, pos *utilities.Point) []utilities.Point {
	ret := []utilities.Point{}

	index := pos.X + pos.Y*world.Width
	my_height := world.Tiles[index]

	if pos.X > 0 && my_height >= world.Tiles[pos.X-1+pos.Y*world.Width]-1 {
		ret = append(ret, utilities.Point{X: pos.X - 1, Y: pos.Y})
	}

	if pos.Y > 0 && my_height >= world.Tiles[pos.X+(pos.Y-1)*world.Width]-1 {
		ret = append(ret, utilities.Point{X: pos.X, Y: pos.Y - 1})
	}

	if pos.X < world.Width-1 && my_height >= world.Tiles[pos.X+1+pos.Y*world.Width]-1 {
		ret = append(ret, utilities.Point{X: pos.X + 1, Y: pos.Y})
	}

	if pos.Y < world.Height-1 && my_height >= world.Tiles[pos.X+(pos.Y+1)*world.Width]-1 {
		ret = append(ret, utilities.Point{X: pos.X, Y: pos.Y + 1})
	}

	return ret
}

func DjikstraButFromAllLowestPoints(world *utilities.World) int {
	queue := make(utilities.PriorityQueue, world.Height*world.Width)

	book := make([]*utilities.Item, world.Height*world.Width)

	index := 0
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			priority := math.MaxInt

			if world.Tiles[x+y*world.Width] == 0 {
				priority = 0
			}

			other := utilities.Item{
				Value: DjikstraPoint{
					Position: utilities.Point{X: x, Y: y},
				},
				Priority: priority,
				Index:    index,
			}

			book[x+y*world.Width] = &other
			queue[index] = &other
			index++
		}
	}

	heap.Init(&queue)

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*utilities.Item)
		point := current.Value.(DjikstraPoint)

		if point.Position.X == world.End.X && point.Position.Y == world.End.Y {
			return current.Priority
		}

		neighbours := Neighbours(world, &point.Position)

		for _, np := range neighbours {
			n := book[np.X+np.Y*world.Width]
			nd := n.Value.(DjikstraPoint)

			new_dist := current.Priority + 1
			if n.Priority > new_dist {
				nd.Previous = current

				n.Value = nd
				n.Priority = new_dist
				n.Value = nd

				heap.Fix(&queue, n.Index)
			}
		}
	}

	return math.MaxInt
}

func PartB(world *utilities.World) int {
	return DjikstraButFromAllLowestPoints(world)
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
