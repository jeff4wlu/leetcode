package infra

import "math"

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

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num < 4 {
		return true
	}

	sq := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func GetPrimes(num int) []int {

	if num <= 1 {
		return []int{}
	}

	res := []int{2}

	tmp := make([]int, num+1)
	for i := 3; i <= num; i++ {
		tmp[i] = i
	}

	for {
		i := len(res) - 1
		for j := 1; j*res[i] <= num; j++ {
			tmp[j*res[i]] = 0
		}
		for k := res[i] + 1; k <= num; k++ {
			if tmp[k] != 0 {
				res = append(res, k)
				break
			}
		}
		if i == len(res)-1 {
			break
		}
	}
	return res
}
