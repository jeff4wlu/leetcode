package leet36

func ValidSuduku(sudu [][]int) bool {

	rowMap := make([]map[int]int, 9)
	colMap := make([]map[int]int, 9)
	//boxMap := make([]map[int]int, 9)
	var boxMap [][]map[int]int

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
			if sudu[i][j] == -1 {
				continue
			}

			if _, ok := rowMap[j][sudu[i][j]]; !ok {
				rowMap[j][sudu[i][j]] = 1
			} else {
				return false
			}

			if _, ok := colMap[i][sudu[i][j]]; !ok {
				colMap[i][sudu[i][j]] = 1
			} else {
				return false
			}

			c := j / 3
			r := i / 3
			if _, ok := boxMap[c][r][sudu[i][j]]; !ok {
				boxMap[c][r][sudu[i][j]] = 1
			} else {
				return false
			}

		}
	}

	return true
}
