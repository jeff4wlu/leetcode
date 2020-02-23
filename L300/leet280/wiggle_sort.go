package leet280

func WiggleSort(nums []int) []int {

	n := len(nums)
	if n < 2 {
		return []int{}
	}

	for i := 1; i < n; i++ {
		if i%2 == 0 { //偶数位置
			if nums[i-1] < nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		} else {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
	}

	return nums
}
