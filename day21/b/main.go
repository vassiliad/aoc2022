package main

import (
	"fmt"
	"os"

	"github.com/vassiliad/aoc2022/day21/a/utilities"
)

const (
	// VV: Unknown indicates that the node in question depends on a value that we have not
	// identified thus far (e.g. `humn` OR something that implicitly uses the value of `humn`)
	TaintUnknown = 1
	TaintKnown   = 0
)

type CalculatorWithTaint struct {
	Tainted map[string]int8
	RefBy   map[string][]string

	utilities.Calculator
}

func Compute(calc *utilities.Calculator, res string) int {
	// fmt.Printf("%+v\n", *calc)

	remaining := []string{res}
	for len(remaining) > 0 {
		res = remaining[len(remaining)-1]

		if _, ok := calc.Registers[res]; ok {
			remaining = remaining[:len(remaining)-1]
			continue
		} else {
			// VV: We have not computed the value YET. If we know the values of operands, compute the value now
			// else, just ask future loops to compute the operands for us
			expr, ok := calc.Expressions[res]

			if !ok {
				panic(fmt.Sprintf("Could not find expression for %s", res))
			}

			missing := 0
			to_add := []string{}
			if _, ok = calc.Registers[expr.Left]; !ok {
				missing++
				to_add = append(to_add, expr.Left)
			}

			if _, ok = calc.Registers[expr.Right]; !ok {
				missing++
				to_add = append(to_add, expr.Right)
			}

			if missing == 0 {
				value, _ := calc.Compute(&expr)
				calc.Registers[res] = value
				remaining = remaining[0 : len(remaining)-1]
			} else {
				remaining = append(remaining, to_add...)
			}
		}
	}

	return calc.Registers[res]
}

func (c *CalculatorWithTaint) Taint(node string, v int8) {
	remaining := []string{node}

	for len(remaining) > 0 {
		node := remaining[len(remaining)-1]
		c.Tainted[node] = v

		remaining = remaining[:len(remaining)-1]

		if ref_by, ok := c.RefBy[node]; ok {
			remaining = append(remaining, ref_by...)
		} else if node != "root" {
			panic(fmt.Sprintf("node %s is not referenced by anyone", node))
		}
	}
}

// Removes Expressions for which we know the value they evaluate to
func (c *CalculatorWithTaint) Minimize() {
	started_with := len(c.Expressions)
	for out, _ := range c.Expressions {
		if c.Tainted[out] == TaintKnown {
			// VV: This will populate c.Calculator.Registers[out] (may already exist)
			Compute(&c.Calculator, out)
			delete(c.Expressions, out)
		}
	}

	fmt.Println("Removed", started_with-len(c.Expressions), "out of", started_with, "expressions")
}

func (c *CalculatorWithTaint) AdaptToValue(expr utilities.Expression, value_parent int) {
	fmt.Printf("Adapting %s=%+v to %d\n", expr.Out, expr, value_parent)

	if c.Tainted[expr.Left]+c.Tainted[expr.Right] == 2 {
		panic(fmt.Sprintf("Both %s and %s are tainted", expr.Left, expr.Right))
	}

	known := expr.Left
	unknown := expr.Right
	if c.Tainted[known] == TaintUnknown {
		known, unknown = unknown, known
	}

	value_sibling := Compute(&c.Calculator, known)
	value_missing := 0

	if expr.Op == utilities.OpAdd {
		value_missing = value_parent - value_sibling
	} else if expr.Op == utilities.OpSub {
		if expr.Left == unknown {
			value_missing = value_sibling + value_parent
		} else {
			value_missing = value_sibling - value_parent
		}
	} else if expr.Op == utilities.OpMult {
		value_missing = value_parent / value_sibling
	} else if expr.Op == utilities.OpDiv {
		if expr.Left == unknown {
			value_missing = value_sibling * value_parent
		} else {
			value_missing = value_sibling / value_parent
		}
	} else {
		panic(fmt.Sprintf("Expression %+v has invalid operator", expr))
	}

	if unknown == "humn" {
		c.Registers[unknown] = value_missing
		return
	}

	c.Taint(unknown, TaintKnown)

	if target, ok := c.Expressions[unknown]; ok {
		c.AdaptToValue(target, value_missing)
	} else {
		panic(fmt.Sprintf("Unknown target %s does not map to an expression", unknown))
	}

}

// If either Left or Right are not tainted (i.e they have a value)
// the method will tweak the other one to have the same value
func (c *CalculatorWithTaint) Solve(left, right string) {
	if c.Tainted[left]+c.Tainted[right] == 2 {
		panic(fmt.Sprintf("Both %s and %s are tainted", left, right))
	}

	known := left
	unknown := right
	if c.Tainted[known] == TaintUnknown {
		known, unknown = unknown, known
	}

	value := Compute(&c.Calculator, known)
	expr, ok := c.Expressions[unknown]

	if !ok {
		panic(fmt.Sprintf("%s does not map to an expression", unknown))
	}

	fmt.Printf("%s value is %d will solve %s=%+v\n", known, value, unknown, expr)
	c.AdaptToValue(expr, value)
}

func PartB(calculator *utilities.Calculator) int {
	calc := CalculatorWithTaint{
		Tainted:    map[string]int8{},
		RefBy:      map[string][]string{},
		Calculator: *calculator,
	}

	for out, expr := range calc.Expressions {
		left := calc.RefBy[expr.Left]
		right := calc.RefBy[expr.Right]

		left = append(left, out)
		right = append(right, out)

		calc.RefBy[expr.Left] = left
		calc.RefBy[expr.Right] = right
	}

	calc.Taint("humn", TaintUnknown)
	delete(calc.Registers, "humn")
	calc.Minimize()

	root := calc.Expressions["root"]

	fmt.Println("Left", root.Left, "is", calc.Tainted[root.Left])
	fmt.Println("Right", root.Right, "is", calc.Tainted[root.Right])

	calc.Solve(root.Left, root.Right)

	return calc.Registers["humn"]
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
