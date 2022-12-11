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
	fmt.Printf("Monkey %d: inspected %d items\n", m.MonkeyLogic.Id, m.TotalInspections)
}

func (m *Monkey) MonkeyBusiness(monkeys []*Monkey, least_common_multiple int) {
	m.TotalInspections += len(m.Items)

	for len(m.Items) > 0 {
		item := m.Items[0]
		m.Items = m.Items[1:]

		if m.MonkeyLogic.OperationAdd != nil {
			item = m.MonkeyLogic.OperationAdd.Calc(item)
		} else {
			item = m.MonkeyLogic.OperationMult.Calc(item)
		}

		dest := monkeys[m.MonkeyLogic.ThrowToMonkeyFalse]
		if item%m.MonkeyLogic.TestDivisibleBy == 0 {
			dest = monkeys[m.MonkeyLogic.ThrowToMonkeyTrue]
		}

		item = item % least_common_multiple

		dest.Items = append(dest.Items, item)
	}

}

func ComputeGreatestCommonDivisor(n1, n2 int) int {
	// VV: wikipedia is great
	for {
		n1, n2 = n2, n1%n2

		if n2 == 0 {
			return n1
		}
	}
}

func ComputeLeastCommonMultiple(numbers []int) int {
	lcm := utilities.AbsInt(numbers[0]*numbers[1]) / ComputeGreatestCommonDivisor(numbers[0], numbers[1])

	for _, v := range numbers[2:] {
		lcm = utilities.AbsInt(v*lcm) / ComputeGreatestCommonDivisor(v, lcm)
	}

	return lcm
}

func PartB(monkey_logic []*utilities.MonkeyLogic) int {
	const rounds = 10000

	monkeys := make([]*Monkey, len(monkey_logic))

	for i, ml := range monkey_logic {
		monkeys[i] = new(Monkey)
		monkeys[i].MonkeyLogic = ml
		monkeys[i].Items = ml.StartingItems
	}

	log := utilities.SetupLogger()

	dividers := make([]int, len(monkeys))
	for i, m := range monkey_logic {
		dividers[i] = m.TestDivisibleBy
	}

	/*VV: If there were 2 monkeys with identical ModuloChecks
	We would just modulo the item value with the modulo-check and there would
	be no exploding numbers.

	However, we can have multiple monkeys and each one of them may use a different
	number to perform its modulo-check. Moreover, all of them may end up
	inspecting the same item in 1 round (at most once). The LCM enables us to
	decrease the magnitude of an item with a value which is guaranteed to be
	divisible by *all* numbers that monkeys use in modulo-tests. Therefore
	"item" and "item % LCM" should be handled exactly the same way.

	For sufficiently small amount of monkeys, and "small" modulo check numbers
	using the LCM is good-enough to avoid the explosion of the item values.
	*/
	least_common_multiple := ComputeLeastCommonMultiple(dividers)

	for i := 1; i <= rounds; i++ {
		for _, m := range monkeys {
			m.MonkeyBusiness(monkeys, least_common_multiple)
		}

		if i == 1 || i == 20 || i%1000 == 0 {
			log.Println("After Round:", i)
			for _, m := range monkeys {
				m.Report()
			}

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
	sol := PartB(input)
	logger.Println("Solution is", sol)
}
