package leet53

import "leetcode/infra"

//暴力破解,复杂度为O(n2)
func MaxSubarr1(in []int) int {

	n := len(in)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return in[0]
	}

	max := infra.INT32_MIN
	for i := 0; i < n; i++ {
		max = infra.IntMax(max, in[i])
		tmp := in[i]
		for j := i + 1; j < n; j++ {
			tmp += in[j]
			max = infra.IntMax(max, tmp)
		}
	}

	return max
}

//DP,复杂度是O(n)
func MaxSubarr2(in []int) int {

	n := len(in)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return in[0]
	}

	max := infra.INT32_MIN
	sum := 0
	for i := 0; i < n; i++ {
		if sum < 0 {
			sum = 0
		}
		sum += in[i]
		max = infra.IntMax(sum, max)

	}

	return max
}

//分治，列举出所有可能性，逐步分解问题。最后汇总。
func MaxSubarr3(in []int) int {
	return helper(in)
}

func helper(in []int) int {

	n := len(in)
	if n == 0 {
		return infra.INT32_MIN //注意这里不是0
	}
	if n == 1 {
		return in[0]
	}

	mid := (n+1)/2 - 1 //求中间元素的索引值

	left := helper(in[:mid])
	right := helper(in[mid+1:])

	max := in[mid]
	sum := in[mid]
	for i := mid - 1; i >= 0; i-- {
		sum += in[i]
		max = infra.IntMax(sum, max)
	}
	sum = max //此句很关键
	for i := mid + 1; i < n; i++ {
		sum += in[i]
		max = infra.IntMax(sum, max)
	}
	max = infra.IntMax(left, max)
	max = infra.IntMax(right, max)

	return max

}
