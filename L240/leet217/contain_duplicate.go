package leet217

func ContainDuplicate(nums []int) bool {

	n := len(nums)
	if n <= 1 {
		return false
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
