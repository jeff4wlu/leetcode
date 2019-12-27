package leet53

import "leetcode/infra"

//DP
func MaxSubarr(nums []int) int {

	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	ans := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		ans = infra.IntMax(ans, dp[i])

	}

	return ans
}

//分治
func MaxSubarr2(nums []int) int {

	return findMaxSub(nums, 0, len(nums)-1)
}

func findMaxSub(nums []int, low, high int) int {

	if low == high {
		return nums[low]
	}

	mid := (high + low) / 2

	leftMax := findMaxSub(nums, low, mid)
	rightMax := findMaxSub(nums, mid+1, high)
	crossMax := findMaxCrossSub(nums, low, mid, high)

	if leftMax >= rightMax && leftMax >= crossMax {
		return leftMax
	} else if rightMax >= leftMax && rightMax >= crossMax {
		return rightMax
	}

	return crossMax
}

func findMaxCrossSub(nums []int, low, mid, high int) int {

	maxLeft := infra.INT32_MIN
	maxRight := infra.INT32_MIN
	var sum int

	for i := mid; i >= low; i-- {
		sum += nums[i]
		if sum > maxLeft {
			maxLeft = sum
			//left = i
		}
	}

	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += nums[i]
		if sum > maxRight {
			maxRight = sum
			//right = i
		}
	}

	return maxLeft + maxRight
}
