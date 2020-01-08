package infra

//-1代表找不到
func BinarySearch(in []int, target int) int {

	n := len(in)
	if n == 0 {
		return -1
	}

	if n == 1 {
		if in[0] == target {
			return 0
		}
		return -1
	}

	l, r := 0, n-1
	var mid int
	for ; l < r; mid = (r+2)/2 - 1 { //这个是索引 {
		if in[mid] == target {
			return mid
		}

		if in[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
