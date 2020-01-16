package leet118

func PascalTriangle(n int) [][]int {

	if n <= 0 {
		return nil
	}

	res := [][]int{}

	for i := 0; i < n; i++ {
		tmp := make([]int, i+1)
		res = append(res, tmp)
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
		if i > 1 {
			for j := 1; j < len(res[i])-1; j++ {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}

	return res

}
