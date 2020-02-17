package leet275

//折半查找和双指针
func HIndex(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2
		tmp := n - mid
		if tmp == nums[mid] {
			return tmp
		}
		if tmp > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left > right {
		return nums[left-1]
	}

	return nums[left]
}
