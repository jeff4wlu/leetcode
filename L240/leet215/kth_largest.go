package leet215

import "leetcode/infra"

//数组内元素不重复，参考快速排序
func KthLargest(nums []int, k int) int {
	return dfs(nums, k)
}

func dfs(nums []int, k int) int {
	n := len(nums)
	if n < 1 || k == 0 || k > n {
		return infra.INT32_MIN
	}

	mid := 0 + (n-1)/2 //0base

	more, less := []int{}, []int{}

	for i := 0; i < mid; i++ {
		if nums[i] > nums[mid] {
			more = append(more, nums[i])
		} else {
			less = append(less, nums[i])
		}
	}
	for i := n - 1; i > mid; i-- {
		if nums[i] >= nums[mid] {
			more = append(more, nums[i])
		} else {
			less = append(less, nums[i])
		}
	}
	numsOfMore := len(more)
	numsOfLess := len(less)
	if numsOfLess == 0 && numsOfMore == 0 || numsOfMore+1 == k {
		return nums[mid]
	}

	if numsOfMore < k {
		return dfs(less, k-numsOfMore-1)
	}
	return dfs(more, k)

}
