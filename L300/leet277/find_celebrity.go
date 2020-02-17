package leet277

//使用图的出度和入度，O(n2)效率不高
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

//根据题意总结出隐含条件，最多只有一个明星；明星不认识任何人。
//效率为O(n)
func FindCelebrity1(relations [][]int) int {

	n := len(relations)
	if n < 2 {
		return -1
	}

	var k int
	for i := 1; i < n; i++ {
		if relations[i][k] == 0 { //i不认识k
			k = i
		}
	}

	for i := 0; i < n; i++ {
		if k != i && (relations[k][i] == 1 || relations[i][k] == 0) {
			return -1
		}
	}

	return k

}
