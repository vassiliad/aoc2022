package utilities

func MaxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

func AbsInt(val int) int {
	return MaxInt(val, -val)
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
	lcm := AbsInt(numbers[0]*numbers[1]) / ComputeGreatestCommonDivisor(numbers[0], numbers[1])

	for _, v := range numbers[2:] {
		lcm = AbsInt(v*lcm) / ComputeGreatestCommonDivisor(v, lcm)
	}

	return lcm
}
