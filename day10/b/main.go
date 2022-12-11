package main

import (
	"os"
	"strings"

	utilities "github.com/vassiliad/aoc2022/day10/a/utilities"
)

type CPU struct {
	Cycle     int
	Registers map[string]int
	CRTBuffer string
}

func NewCPU() CPU {
	return CPU{Registers: map[string]int{"X": 1}}
}

func AddCycleStart(add *utilities.OpCodeAdd, cpu *CPU) {
	add.Cycle += 1
}

func AddCycleStop(add *utilities.OpCodeAdd, cpu *CPU) bool {
	if add.Cycle == 2 {
		cpu.Registers[add.Register] += add.Immediate
		return true
	}

	return false
}

func NoopCycleStart(noop *utilities.OpCodeNoop, cpu *CPU) {

}

func NoopCycleStop(noop *utilities.OpCodeNoop, cpu *CPU) bool {
	return true
}

func (c *CPU) CycleStart() {
	wrapped := c.Cycle % 40

	if utilities.AbsInt(wrapped-c.Registers["X"]) < 2 {
		c.CRTBuffer += "#"
	} else {
		c.CRTBuffer += "."
	}

	if (c.Cycle+1)%40 == 0 {
		c.CRTBuffer += "\n"
	}
	c.Cycle++
}

func (c *CPU) CycleStop() {

}

func PartB(instructions []any) string {
	cpu := NewCPU()

	for _, instr := range instructions {
		if add, ok := instr.(utilities.OpCodeAdd); ok {
			done := false
			for !done {
				cpu.CycleStart()
				AddCycleStart(&add, &cpu)
				cpu.CycleStop()
				done = AddCycleStop(&add, &cpu)
			}
		} else if noop, ok := instr.(utilities.OpCodeNoop); ok {
			done := false
			for !done {
				cpu.CycleStart()
				NoopCycleStart(&noop, &cpu)
				cpu.CycleStop()
				done = NoopCycleStop(&noop, &cpu)
			}
		} else {
			panic(instr)
		}
	}

	return strings.TrimSpace(cpu.CRTBuffer)
}

func main() {
	logger := utilities.SetupLogger()

	logger.Println("Parse input")
	input, err := utilities.ReadInputFile(os.Args[1])

	if err != nil {
		logger.Fatalln("Run into problems while reading input. Problem", err)
	}
	sol := PartB(input)
	logger.Printf("Solution is\n%s\n", strings.ReplaceAll(sol, ".", " "))
}
