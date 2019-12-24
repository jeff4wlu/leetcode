package leet41

//难点在于限制条件是空间O(1)，所以只能依靠本身数组的下标来判断最小整数
func FirstMissingPositive(nums []int) int {

	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] <= 0 || nums[i] == i+1 { //不移动
			continue
		}
		for {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			if nums[i] > n || nums[i] <= 0 || nums[i] == i+1 { //不移动
				break
			}
		}

	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 { //不移动
			return i + 1
		}

	}
	return n + 1
}
