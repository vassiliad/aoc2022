package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	Name    string
	UID     int
	Rate    int
	FlowsTo BitMask
}

func (v *Valve) GetNeighbourNames(network *Network) []string {
	ret := []string{}

	for i := 0; i < len(network.Valves); i++ {
		if v.FlowsTo.Get(i) {
			name, ok := network.BookReverse[i]
			if !ok {
				panic(i)
			}
			ret = append(ret, name)
		}
	}

	return ret
}

func (v *Valve) GetNeighbourUIDS(network *Network) []int {
	ret := []int{}

	for i := 0; i < len(network.Valves); i++ {
		if v.FlowsTo.Get(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

type Network struct {
	Valves      map[string]*Valve
	BookReverse map[int]string
}

func ReadScanner(scanner *bufio.Scanner) (*Network, error) {
	ret := Network{
		Valves:      make(map[string]*Valve),
		BookReverse: make(map[int]string),
	}

	flows_to := map[string][]string{}

	line_idx := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		rate, err := strconv.Atoi(tokens[4][5 : len(tokens[4])-1])

		if err != nil {
			return &ret, fmt.Errorf("unable to decode rate %s in line %s due to %s", tokens[4], line, err)
		}

		valve := Valve{Name: tokens[1], UID: line_idx, Rate: rate}

		ret.BookReverse[valve.UID] = valve.Name
		ret.Valves[valve.Name] = &valve

		this_flow := []string{}
		for _, other := range tokens[9:] {
			if len(other) > 2 {
				other = other[:len(other)-1]
			}
			this_flow = append(this_flow, other)
		}
		flows_to[valve.Name] = this_flow

		line_idx++
	}

	for name, downstream := range flows_to {
		if valve, ok := ret.Valves[name]; ok {
			for _, d := range downstream {
				if dv, ok := ret.Valves[d]; ok {
					valve.FlowsTo.Set(dv.UID)
				} else {
					return &ret, fmt.Errorf("tried to backpatch valve %s with %s but the later does not exist", name, d)
				}
			}
		} else {
			return &ret, fmt.Errorf("tried to backpatch valve %s but it does not exist", name)
		}
	}

	return &ret, scanner.Err()
}

func ReadString(text string) (*Network, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) (*Network, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
