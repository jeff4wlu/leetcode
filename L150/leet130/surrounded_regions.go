package leet130

func SurroundedRegions(board *[][]byte) {

	row := len(*board)
	col := len((*board)[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (i == 0 || j == 0 || i == row-1 || j == col-1) && (*board)[i][j] == 'o' {
				dfs(board, i, j, row, col)
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (*board)[i][j] == 'o' {
				(*board)[i][j] = 'x'
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (*board)[i][j] == '$' {
				(*board)[i][j] = 'o'
			}
		}
	}

}

func dfs(board *[][]byte, x, y, row, col int) {

	if (*board)[x][y] == 'o' {
		(*board)[x][y] = '$'
		if x < row-1 && (*board)[x+1][y] == 'o' {
			dfs(board, x+1, y, row, col)
		}
		if x > 0 && (*board)[x-1][y] == 'o' {
			dfs(board, x-1, y, row, col)
		}

		if y < col-1 && (*board)[x][y+1] == 'o' {
			dfs(board, x, y+1, row, col)
		}
		if y > 0 && (*board)[x][y-1] == 'o' {
			dfs(board, x, y-1, row, col)
		}

	}

}
