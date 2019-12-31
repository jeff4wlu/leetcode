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
