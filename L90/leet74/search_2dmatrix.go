package leet74

func Search2DMatrix(in [][]int, num int) bool {

	rSize := len(in)
	cSize := len(in[0])

	var left, right, mid, row int

	//2分法找行
	right = rSize - 1

	for left <= right {
		mid = (left + right) / 2
		if num == in[mid][cSize-1] {
			return true
		} else if num > in[mid][cSize-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if num > in[left][cSize-1] {
		row = right
	} else {
		row = right - 1
	}
	if row < 0 {
		row = 0
	}

	left = 0
	right = cSize - 1

	//2分法找列
	for left <= right {
		mid = (left + right) / 2
		if num == in[row][mid] {
			return true
		} else if num > in[row][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
