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
