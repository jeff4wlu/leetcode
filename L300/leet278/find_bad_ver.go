package leet278

//二分法，bad为1，好为0
//1base
func FindBadVer(ver []int) int {

	n := len(ver)

	left, right := 1, n-1

	for left <= right {
		mid := left + (right-left)/2
		if ver[mid] == 1 { //坏的版本
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
