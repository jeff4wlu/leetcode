package leet59

func SpiralMatrix2(num int) [][]int {

	top := 0
	down := num - 1
	left := 0
	right := num - 1

	res := make([][]int, num)
	for i := 0; i < num; i++ {
		res[i] = make([]int, num)
	}

	step := 1
	for {
		for i := left; i <= right; i++ {
			res[top][i] = step
			step++
		}
		top++
		if top > down {
			break
		}

		for i := top; i <= down; i++ {
			res[i][right] = step
			step++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res[down][i] = step
			step++
		}
		down--
		if top > down {
			break
		}

		for i := down; i >= top; i-- {
			res[i][left] = step
			step++
		}
		left++
		if left > right {
			break
		}
	}

	return res

}
