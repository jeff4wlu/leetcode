package leet35

func SearchInsertPos(in []int, target int) int {

	n := len(in)

	var left, right int
	right = n - 1

	for left <= right {
		mid := left + (right-left)/2
		if in[mid] == target {
			return mid
		} else if in[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
