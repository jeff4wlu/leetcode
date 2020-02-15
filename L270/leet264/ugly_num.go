package leet264

import "leetcode/infra"

//思路，使用DP算法。第k个丑数肯定是前面丑数*2、3或5中最小的一个。
func UglyNum(k int) int {

	if k <= 0 {
		return -1
	}
	if k == 1 {
		return 1
	}
	idx2, idx3, idx5 := 1, 1, 1

	var tmp int
	for i := 1; i < k; i++ {
		tmp = infra.IntMin(infra.IntMin(idx2*2, idx3*3), idx5*5)
		if tmp == idx2*2 {
			idx2++
		}
		if tmp == idx3*3 {
			idx3++
		}
		if tmp == idx5*5 {
			idx5++
		}
	}
	return tmp
}
