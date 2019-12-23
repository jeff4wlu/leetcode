package leet36

func ValidSudoku(sudoku [][]int) bool {

	var rowMap, colMap []map[int]int
	var boxMap [][]map[int]int

	rowMap = make([]map[int]int, 9)
	colMap = make([]map[int]int, 9)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[int]int)
		colMap[i] = make(map[int]int)
	}
	boxMap = make([][]map[int]int, 3)
	for i := 0; i < 3; i++ {
		boxMap[i] = make([]map[int]int, 3)
		for j := 0; j < 3; j++ {
			boxMap[i][j] = make(map[int]int)
		}
	}

	//i是列，j是行
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				continue
			}

			if _, ok := rowMap[j][sudoku[i][j]]; !ok {
				rowMap[j][sudoku[i][j]] = 1
			} else {
				return false
			}

			if _, ok := colMap[i][sudoku[i][j]]; !ok {
				colMap[i][sudoku[i][j]] = 1
			} else {
				return false
			}

			c := j / 3
			r := i / 3
			if _, ok := boxMap[c][r][sudoku[i][j]]; !ok {
				boxMap[c][r][sudoku[i][j]] = 1
			} else {
				return false
			}

		}
	}

	return true
}
