package leet69

func Sqrt(num int) int {

	start := 0
	end := num

	for start+1 < end {

		mid := start + (end-start)/2

		if mid*mid == num {
			return mid
		} else if mid*mid < num {
			start = mid
		} else {
			end = mid
		}
	}

	if end*end == num {
		return end
	}

	return start
}
