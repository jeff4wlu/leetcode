package leet153

import "leetcode/infra"

func MinRotatedSortedArr(in []int) int {
	return dfs(in)
}

func dfs(in []int) int {
	n := len(in)

	if n == 0 {
		return infra.INT32_MAX
	}

	if n == 1 {
		return in[0]
	}

	var min int
	mid := (n+1)/2 - 1
	if in[mid] > in[n-1] {
		min = in[0]
		min = infra.IntMin(min, dfs(in[mid+1:]))
	} else {
		min = in[mid]
		min = infra.IntMin(min, dfs(in[:mid]))
	}
	return min
}
