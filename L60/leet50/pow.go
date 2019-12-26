package leet50

//分治
func Pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	//下句是一个优化，减少重复计算。使时间复杂度为log(n)
	half := Pow(x, n/2)

	if n%2 == 0 {
		return half * half
	}

	if n > 0 {
		return x * half * half
	}
	return half * half / x

}

func Pow2(x float64, n int) float64 {

	var m int
	if n < 0 {
		m = -n - 1
	} else {
		m = n
	}

	p := 1.0

	for q := x; m > 0; m /= 2 {
		if m&1 != 0 {
			p *= q
		}
		q *= q
	}
	if n < 0 {
		return 1 / p / x
	}
	return p
}
