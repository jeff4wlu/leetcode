package leet62

//使用递归
func UniquePaths(rows, cols int) (int, [][]string) {

	res := [][]string{}
	path := []string{}
	solve(rows, cols, 0, 1, "right", path, &res)
	solve(rows, cols, 1, 0, "down", path, &res)
	return len(res), res
}

func solve(rows, cols, i, j int, direction string, path []string, res *[][]string) {

	//到达终点
	if i == rows-1 && j == cols-1 {

		path = append(path, direction)
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	//出界
	if i >= rows || j >= cols {
		return
	}

	path = append(path, direction)

	solve(rows, cols, i, j+1, "right", path, res)
	solve(rows, cols, i+1, j, "down", path, res)
}

//TODO:使用DP
//首先找出状态所依靠的变量，此题为坐标i,j。
//其次找出状态转换方程，numsWay[i][j] = numsWay[i-1][j]+numsWay[i][j-1]
func UniquePaths2(rows, cols int) int {

	//初始化一个二维数组，用作记录过程。空间换时间
	numsWay := [][]int{}
	for i := 0; i < rows; i++ {
		r := make([]int, cols)
		numsWay = append(numsWay, r)
	}

	numsWay[0][0] = 1
	for i := 1; i < cols; i++ {
		numsWay[0][i] = 1
	}
	for i := 1; i < rows; i++ {
		numsWay[i][0] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			numsWay[i][j] = numsWay[i-1][j] + numsWay[i][j-1]
		}
	}

	return numsWay[rows-1][cols-1]
}
