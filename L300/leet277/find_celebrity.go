package leet277

//使用图的出度和入度
func FindCelebrity(relations [][]int) int {

	n := len(relations)
	if n < 2 {
		return -1
	}

	in := make([]int, n)
	out := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if relations[i][j] == 1 {
					in[j]++
					out[i]++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i
		}
	}

	return -1

}
