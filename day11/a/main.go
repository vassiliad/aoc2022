package main

import (
	"fmt"
	"os"
	"sort"

	utilities "github.com/vassiliad/aoc2022/day11/a/utilities"
)

type Monkey struct {
	MonkeyLogic      *utilities.MonkeyLogic
	Items            []int
	TotalInspections int
}

func (m *Monkey) Report() {
	fmt.Printf("Monkey %d:", m.MonkeyLogic.Id)
	for _, x := range m.Items {
		fmt.Printf("%d, ", x)
	}
	fmt.Println()
}

func (m *Monkey) MonkeyBusiness(monkeys []*Monkey) {
	m.TotalInspections += len(m.Items)

	for len(m.Items) > 0 {
		item := m.Items[0]
		m.Items = m.Items[1:]

		if m.MonkeyLogic.OperationAdd != nil {
			item = m.MonkeyLogic.OperationAdd.Calc(item)
		} else {
			item = m.MonkeyLogic.OperationMult.Calc(item)
		}

		item = item / 3

		if item%m.MonkeyLogic.TestDivisibleBy == 0 {
			monkeys[m.MonkeyLogic.ThrowToMonkeyTrue].Items = append(monkeys[m.MonkeyLogic.ThrowToMonkeyTrue].Items, item)
		} else {
			monkeys[m.MonkeyLogic.ThrowToMonkeyFalse].Items = append(monkeys[m.MonkeyLogic.ThrowToMonkeyFalse].Items, item)
		}
	}

}

func PartA(monkey_logic []*utilities.MonkeyLogic) int {
	const rounds = 20

	monkeys := make([]*Monkey, len(monkey_logic))

	for i, ml := range monkey_logic {
		monkeys[i] = new(Monkey)
		monkeys[i].MonkeyLogic = ml
		monkeys[i].Items = ml.StartingItems
	}

	log := utilities.SetupLogger()

	for i := 1; i <= rounds; i++ {
		for _, m := range monkeys {
			m.MonkeyBusiness(monkeys)
		}

		log.Println("After Round:", i)
		for _, m := range monkeys {
			m.Report()
		}
	}

	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.TotalInspections
	}

	sort.Ints(inspections)

	return inspections[len(monkeys)-1] * inspections[len(monkeys)-2]
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
