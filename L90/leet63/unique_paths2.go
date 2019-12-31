package leet63

func UniquePaths2(grid [][]int) (int, [][]string) {
	res := [][]string{}
	path := []string{}
	rows := len(grid)
	cols := len(grid[0])
	solve(grid, 0, 1, rows, cols, "right", path, &res)
	solve(grid, 1, 0, rows, cols, "down", path, &res)
	return len(res), res
}

func solve(grid [][]int, i, j, rows, cols int, direction string, path []string, res *[][]string) {

	//到达终点
	if i == rows-1 && j == cols-1 {

		path = append(path, direction)
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	//出界
	if i >= rows || j >= cols || grid[i][j] == 1 {
		return
	}

	path = append(path, direction)

	solve(grid, i, j+1, rows, cols, "right", path, res)
	solve(grid, i+1, j, rows, cols, "down", path, res)
}

func UniquePaths22(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	//初始化一个二维数组，用作记录过程。空间换时间
	numsWay := [][]int{}
	for i := 0; i < rows; i++ {
		r := make([]int, cols)
		numsWay = append(numsWay, r)
	}

	numsWay[0][0] = 1
	for i := 1; i < cols; i++ {
		if grid[0][i] == 1 {
			for j := i; j < cols; j++ {
				numsWay[0][j] = 0
			}
			break
		} else {
			numsWay[0][i] = 1
		}
	}
	for i := 1; i < rows; i++ {
		if grid[i][0] == 1 {
			for j := i; j < cols; j++ {
				numsWay[j][0] = 0
			}
			break
		} else {
			numsWay[i][0] = 1
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if grid[i][j] == 1 {
				numsWay[i][j] = 0
			} else {
				numsWay[i][j] = numsWay[i-1][j] + numsWay[i][j-1]
			}
		}
	}

	return numsWay[rows-1][cols-1]
}
