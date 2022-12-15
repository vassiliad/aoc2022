package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day15/a/utilities"
)

func CountDefinitelyNotBeacon(pairs []utilities.Pair, row, search_max int) (int, bool) {
	dummy := utilities.Point{Y: row}

	const frequency_factor = 4000000

	for x := 0; x <= search_max; x++ {
		dummy.X = x

		is_beacon := true
		for _, p := range pairs {
			distance := dummy.Distance(&p.Sensor)
			if distance <= p.Distance {
				is_beacon = false

				dx := utilities.AbsInt(x - p.Sensor.X)
				dy := utilities.AbsInt(row - p.Sensor.Y)

				delta := 1
				if p.Sensor.X > x {
					delta = 2*p.Distance - dy*2 - 1
				} else if p.Sensor.X <= x {
					delta = p.Distance - dx - dy
				}

				x += delta
				break
			}
		}

		if is_beacon {
			return x*frequency_factor + row, true
		}

	}

	return -1, false
}

func PartB(pairs []utilities.Pair, search_max_row int) int {
	for y := 0; y <= search_max_row; y++ {
		freq, found := CountDefinitelyNotBeacon(pairs, y, search_max_row)

		if found {
			return freq
		}
	}

	panic("Could not find Distress Beacon")
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartB(input, 4000000)
	logger.Println("Solution is", sol)
}
