package leet33

//注意算法要求log(n)的时间
func SearchRotatedSortedArr(in []int, target int) int {

	n := len(in)

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if target == in[mid] {
			return mid
		}
		//右边半段
		if in[mid] < in[right] {
			if target > in[mid] && target <= in[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if in[left] <= target && target < in[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1

}
