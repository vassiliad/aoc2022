package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Resource int

const (
	ResOre         Resource = 0
	ResClay        Resource = 1
	ResObsidian    Resource = 2
	ResGeode       Resource = 3
	TotalResources int      = 4
)

var ResOrder []Resource = []Resource{ResOre, ResClay, ResObsidian, ResGeode}

type ManyParts [TotalResources]int

type Recipe struct {
	Output       Resource
	CostOre      int
	CostClay     int
	CostObsidian int
}

func (r *Recipe) CanProduce(inventory *ManyParts) bool {
	return inventory[ResOre] >= r.CostOre &&
		inventory[ResClay] >= r.CostClay &&
		inventory[ResObsidian] >= r.CostObsidian
}

func (r *Recipe) Pay(inventory *ManyParts) {
	if !r.CanProduce(inventory) {
		panic(r)
	}

	inventory[ResOre] -= r.CostOre
	inventory[ResClay] -= r.CostClay
	inventory[ResObsidian] -= r.CostObsidian
}

type Blueprint struct {
	Robots [TotalResources]Recipe
	ID     int
}

func ParseResource(res string) (Resource, error) {
	if res == "ore" {
		return ResOre, nil
	} else if res == "clay" {
		return ResClay, nil
	} else if res == "obsidian" {
		return ResObsidian, nil
	} else if res == "geode" {
		return ResGeode, nil
	} else {
		return 0, fmt.Errorf("unable to decode Resource from %s", res)
	}
}

func (r *Recipe) parseLine(line string) error {
	tokens := strings.Split(line, " ")
	prod, err := ParseResource(tokens[0])

	if err != nil {
		return fmt.Errorf("could not decode Robot.Produce because of %s", err)
	}

	r.Output = prod

	components := tokens[3:]

	if (len(components)-2)%3 != 0 {
		return fmt.Errorf("unexpected components %+v", components)
	}

	for i := 0; i < len(components); i += 3 {
		amount, err := strconv.Atoi(components[i])

		if err != nil {
			return fmt.Errorf("cannot decode component %s due to %s", components[i], err)
		}
		res_type, err := ParseResource(components[i+1])
		if err != nil {
			return fmt.Errorf("could not decode component because of %s", err)
		}

		if res_type == ResOre {
			r.CostOre = amount
		} else if res_type == ResClay {
			r.CostClay = amount
		} else if res_type == ResObsidian {
			r.CostObsidian = amount
		} else {
			return fmt.Errorf("unknown component %v", res_type)
		}
	}

	return nil
}

func ReadScanner(scanner *bufio.Scanner) ([]Blueprint, error) {
	ret := []Blueprint{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		bp := Blueprint{}

		parts := strings.Split(line, ":")

		the_id := strings.Split(parts[0], " ")[1]
		bp_id, err := strconv.Atoi(the_id)

		if err != nil {
			return ret, fmt.Errorf("unable to decode Blueprint id %s due to %s", the_id, err)
		}

		bp.ID = bp_id

		for _, recipe := range strings.Split(parts[1], " Each ") {
			recipe = strings.TrimSpace(recipe)
			if len(recipe) == 0 {
				continue
			}
			robot := Recipe{}
			err = robot.parseLine(recipe[:len(recipe)-1])
			if err != nil {
				return ret, fmt.Errorf("unable to decode Robot due to %s", err)
			}
			bp.Robots[robot.Output] = robot
		}

		ret = append(ret, bp)
	}

	return ret, scanner.Err()
}

func ReadString(text string) ([]Blueprint, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))

	return ReadScanner(scanner)
}

func ReadInputFile(path string) ([]Blueprint, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	return ReadScanner(scanner)
}
