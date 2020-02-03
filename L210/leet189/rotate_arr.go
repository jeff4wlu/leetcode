package leet189

func RotateArr(nums []int, k int) []int {

	n := len(nums)
	if n <= 1 {
		return nums
	}
	for i := 0; i < k; i++ {
		for j := n - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}
