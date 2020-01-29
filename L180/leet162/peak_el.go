package leet162

func PeakEl(in []int) int {
	n := len(in)

	left, right := 0, n-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if in[mid] > in[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right

}
