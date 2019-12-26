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

//fast power
func Pow2(x float64, n int) float64 {

	var m int
	if n < 0 {
		//如果是n负数，要避免n=-2147483648溢出
		m = -n - 1
	} else {
		m = n
	}

	p := 1.0

	//十进制变二进制时，即是十进制数不断除以2，得到二进制数的各位（先低位）
	for q := x; m > 0; m /= 2 {
		if m&1 != 0 { //判断此权位是否为1，例如4的二进制是100，其最高位是1
			p *= q
		}
		q *= q //m每次除2，q就要变为平方
	}
	if n < 0 {
		return 1 / p / x
	}
	return p
}
