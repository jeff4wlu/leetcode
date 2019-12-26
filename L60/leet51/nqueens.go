package leet51

import "math"

func NQueens(n int) [][]int {
	solution := make([]int, n)
	for i := 0; i < n; i++ {
		solution[i] = -1
	}
	res := [][]int{}
	solve(solution, 0, n, &res)
	return res
}

func solve(solution []int, curRow, size int, res *[][]int) {

	if curRow == size {
		//这里必须重新拷贝一份solution，否则会被后面的覆盖
		tmp := make([]int, size)
		copy(tmp, solution)
		*res = append(*res, tmp)
		return
	}

	//每一列都尝试一下是否可行，求出列.穷举所有的可能性
	for i := 0; i < size; i++ {

		if isValid(solution, curRow, i) {
			solution[curRow] = i //把queen放在i列

			solve(solution, curRow+1, size, res)
		}
		solution[curRow] = -1
	}

}

func isValid(solution []int, curRow, curCol int) bool {

	for i := 0; i < curRow; i++ {
		if solution[i] == -1 {
			continue
		}

		if solution[i] == curCol || math.Abs(float64(i-curRow)) == math.Abs(float64(solution[i]-curCol)) {
			return false
		}
	}
	return true
}
