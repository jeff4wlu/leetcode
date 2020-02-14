package infra

//true，从大到小；
func BubbleSort(nums []int, isBig bool) []int {
	n := len(nums)

	if n >= 2 {
		for i := 0; i < n-1; i++ {
			for j := 1; j < n-i; j++ {
				if !isBig {
					if nums[j] < nums[j-1] {
						nums[j], nums[j-1] = nums[j-1], nums[j]
					}
				} else {
					if nums[j] > nums[j-1] {
						nums[j], nums[j-1] = nums[j-1], nums[j]
					}
				}

			}
		}
	}

	return nums
}
