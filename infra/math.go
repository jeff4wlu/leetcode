package infra

const (
	INT32_MAX = 2147483647
	INT32_MIN = -2147483648
)

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
