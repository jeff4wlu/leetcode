package leet240

//考察折半法找区间
func SearchMatrix(m [][]int, num int) bool {

	row := len(m)
	col := len(m[0])

	var rowLeft, colLeft int
	rowRight, colRight := row-1, col-1
	rowl, rowr, coll, colr := 0, row-1, 0, col-1

	//，再求大端。left是小端，right是大端
	for num <= m[rowr][coll] || num <= m[rowl][colr] {
		//先求小端;区间在小于left大于right
		for rowLeft <= rowRight {
			mid := rowLeft + (rowRight-rowLeft)/2
			if m[mid][colLeft] == num {
				return true
			} else if m[mid][colLeft] < num {
				rowLeft = mid + 1
			} else {
				rowRight = mid - 1
			}
		}
		if rowLeft < row {
			rowr = rowLeft - 1
			rowLeft = rowl
			rowRight = rowr
			row = rowr - rowl + 1
		}

		for colLeft <= colRight {
			mid := colLeft + (colRight-colLeft)/2
			if m[rowLeft][mid] == num {
				return true
			} else if m[rowLeft][mid] < num {
				colLeft = mid + 1
			} else {
				colRight = mid - 1
			}

		}
		if colLeft < row {
			colr = colLeft - 1
			colLeft = coll
			colRight = colr
			col = colr - coll + 1
		}

		//再求大端
		for rowLeft <= rowRight {
			mid := rowLeft + (rowRight-rowLeft)/2
			if m[mid][colRight] == num {
				return true
			} else if m[mid][colRight] < num {
				rowLeft = mid + 1
			} else {
				rowRight = mid - 1
			}
		}
		if rowLeft < row {
			rowr = rowLeft - 1
			rowLeft = rowl
			rowRight = rowr
			row = rowr - rowl + 1
		}

		for colLeft <= colRight {
			mid := colLeft + (colRight-colLeft)/2
			if m[rowRight][mid] == num {
				return true
			} else if m[rowRight][mid] < num {
				colLeft = mid + 1
			} else {
				colRight = mid - 1
			}

		}
		if colLeft < row {
			colr = colLeft - 1
			colLeft = coll
			colRight = colr
			col = colr - coll + 1
		}
	}

	return false
}
