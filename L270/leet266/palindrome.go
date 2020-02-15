package leet266

//题目要求是字符串，所以我们假设是128个值的非扩展ascii码
func Palindrome(word string) bool {

	n := len(word)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	var high, low uint64 //128个字符的ascii码

	//求出mask；取出具体那bit；与mask运算
	for i := 0; i < n; i++ {
		var mask uint64
		k := uint(word[i])
		if k > 63 {
			mask = 1 << (k - 64)

			tmp := (high & mask) >> (k - 64)
			if tmp == 0 {
				high |= mask
			} else {
				high &= ^mask
			}
		} else {
			mask = 1 << k
			tmp := (low & mask) >> k
			if tmp == 0 {
				low |= mask
			} else {
				low &= ^mask
			}
		}
	}

	var count int
	for i := 0; i < 64; i++ {
		if (low&uint64(1<<uint(i)))>>uint(i) == 1 {
			count++
		}
		if (high&uint64(1<<uint(i)))>>uint(i) == 1 {
			count++
		}
	}
	if count < 2 {
		return true
	}
	return false
}
