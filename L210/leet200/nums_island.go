package leet200

//dfs深度搜索
func NumsIsland(m []string) int {

	row := len(m)
	col := len(m[0])

	bmap := [][]int{}

	for i := 0; i < row; i++ {
		tmp := []int{}
		for j := 0; j < col; j++ {
			var a int
			if m[i][j] != '0' {
				a = 1
			}

			tmp = append(tmp, a)
		}
		bmap = append(bmap, tmp)
	}

	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if bmap[i][j] == 1 {
				dfs(&bmap, i, j)
				res++
			}
		}
	}

	return res
}

func dfs(m *[][]int, x, y int) {

	row := len(*m)
	col := len((*m)[0])
	(*m)[x][y] = -1

	//上
	if x-1 >= 0 && (*m)[x-1][y] == 1 {
		dfs(m, x-1, y)
	}
	//下
	if x+1 < row && (*m)[x+1][y] == 1 {
		dfs(m, x+1, y)
	}
	//左
	if y-1 >= 0 && (*m)[x][y-1] == 1 {
		dfs(m, x, y-1)
	}

	//右
	if y+1 < col && (*m)[x][y+1] == 1 {
		dfs(m, x, y+1)
	}

}
