package leet9

// x/位数取左边第一个，x%10取右边第一个
// x%位数去掉最右边，x/10去掉最左边
func PalindromeNum(x int) bool {

	if x < 0 {
		return false
	}

	div := 1

	//求出x的位数
	for x/div >= 10 {
		div *= 10
	}

	for x > 0 {

		left := x / div
		right := x % 10
		if left != right {
			return false
		}
		x = (x % div) / 10
		div /= 100 //去掉两位了已经

	}

	return true

}
