package main

import (
	"os"

	utilities "github.com/vassiliad/aoc2022/day15/a/utilities"
)

func CountDefinitelyNotBeacon(pairs []utilities.Pair, row, left, right int) int {
	not_beacons := 0

	dummy := utilities.Point{Y: row}

	existing := map[int]int8{}

	for _, p := range pairs {
		if p.Beacon.Y == row {
			existing[p.Beacon.X] = 0
		}
	}

	for x := left; x <= right; x++ {
		if _, ok := existing[x]; ok {
			continue
		}

		dummy.X = x

		for _, p := range pairs {
			if dummy.Distance(&p.Sensor) <= p.Distance {
				not_beacons++
				break
			}
		}
	}

	return not_beacons
}

func PartA(pairs []utilities.Pair, row int) int {
	left := utilities.MinInt(pairs[0].Beacon.X-pairs[0].Distance, pairs[0].Sensor.X-pairs[0].Distance)
	right := utilities.MaxInt(pairs[0].Beacon.X+pairs[0].Distance, pairs[0].Sensor.X+pairs[0].Distance)

	for _, p := range pairs[1:] {
		l := utilities.MinInt(p.Beacon.X-p.Distance, p.Sensor.X-p.Distance)
		r := utilities.MaxInt(p.Beacon.X+p.Distance, p.Sensor.X+p.Distance)

		left, right = utilities.MinInt(l, left), utilities.MaxInt(r, right)
	}

	return CountDefinitelyNotBeacon(pairs, row, left, right)
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartA(input, 2000000)
	logger.Println("Solution is", sol)
}
