package leet231

//本思路最直接，但题目要求时间和空间都要是常数复杂度
func PowerTwo(num int) bool {

	if num == 1 {
		return true
	}

	var remainder int
	quotient := num

	for quotient >= 2 && remainder == 0 {
		remainder = quotient % 2 //这个必须在前面，否则quotient被改动
		quotient = quotient / 2
	}

	if remainder != 0 {
		return false
	}

	return true
}

//使用bit操作，分析2的次方的bit特点,低位全是0
func PowerTwo1(num int) bool {

	var cnt int

	for num > 0 {
		cnt += num & 1
		num >>= 1
	}

	return cnt == 1
}
