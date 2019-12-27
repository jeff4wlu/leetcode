package leet54

func SpiralMatrix(matrix [][]int) []int {

	rowStart := 0
	rowEnd := len(matrix) - 1
	colStart := 0
	colEnd := len(matrix[0]) - 1
	res := []int{}

	for {
		//right
		for i := colStart; i <= colEnd; i++ {
			res = append(res, matrix[rowStart][i])
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}

		//down
		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--
		if colStart > colEnd {
			break
		}

		//left
		for i := colEnd; i >= colStart; i-- {
			res = append(res, matrix[rowEnd][i])
		}
		rowEnd--
		if rowStart > rowEnd {
			break
		}

		//up
		for i := rowEnd; i >= rowStart; i-- {
			res = append(res, matrix[i][colStart])
		}
		colStart++
		if rowStart > rowEnd {
			break
		}

	}

	return res

}
