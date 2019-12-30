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

//阶乘
func Fatorial(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}
