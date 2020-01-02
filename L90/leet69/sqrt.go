package leet69

func Sqrt(num int) int {

	var mid int
	end := num + 1
	start := 1
	for end > start {

		mid = start + (end-1)/2
		if mid*mid > num {
			end = mid
		} else {
			start = mid + 1
		}

	}
	return start - 1
}
