package leet85

import "leetcode/L90/leet84"

func MaxRect(in [][]int) int {

	row := len(in)
	col := len(in[0])

	var max int

	for i := 0; i < row; i++ {

		intArr := make([]int, col)

		for j := 0; j < col; j++ {
			for k := 0; k <= i; k++ {
				if in[i-k][j] == 0 {
					break
				}
				intArr[j]++

			}
		}

		tmp := leet84.LargestRect1(intArr)
		if max < tmp {
			max = tmp
		}

	}

	return max

}
