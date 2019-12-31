package leet64

import "leetcode/infra"

//DP算法
func MinPathSum(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	minSum := [][]int{}
	for i := 0; i < rows; i++ {
		r := make([]int, cols)
		minSum = append(minSum, r)
	}

	minSum[0][0] = grid[0][0]

	for i := 1; i < cols; i++ {
		minSum[0][i] = minSum[0][i-1] + grid[0][i]
	}
	for i := 1; i < rows; i++ {
		minSum[i][0] = minSum[i-1][0] + grid[i][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			minSum[i][j] = infra.IntMin(minSum[i-1][j], minSum[i][j-1]) + grid[i][j]
		}
	}

	return minSum[rows-1][cols-1]
}
