package base

const INT32_MAX = 2147483647

func IntMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
