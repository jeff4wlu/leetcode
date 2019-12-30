package leet62

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
