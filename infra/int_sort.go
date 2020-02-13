package infra

func BubbleSort(nums []int) []int {
	n := len(nums)

	if n >= 2 {
		for i := 0; i < n-1; i++ {
			for j := 1; j < n-i; j++ {
				if nums[j] < nums[j-1] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				}
			}
		}
	}

	return nums
}
