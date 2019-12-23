package leet37

import "leetcode/L60/leet36"

func SudokuSolver(sudoku [][]int, i, j int) bool {

	//i是列，i等于9即是整个框走完
	if i == 9 {
		return true
	}
	//j是行，j大于等于9即是要换列了
	if j >= 9 {
		return SudokuSolver(sudoku, i+1, 0)
	}

	if sudoku[i][j] != 0 {
		return SudokuSolver(sudoku, i, j+1)
	}

	for k := 1; k <= 9; k++ {
		sudoku[i][j] = k
		if !leet36.ValidSudoku(sudoku) {
			sudoku[i][j] = 0
			continue
		}
		if SudokuSolver(sudoku, i, j+1) {
			return true
		}
		sudoku[i][j] = 0
	}

	return false
}
