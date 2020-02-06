package leet216

//k个数的和是n。
func CombinationSum3(k, n int) [][]int {

	res := [][]int{}
	path := []int{}
	dfs(1, k, n, path, &res)
	return res
}

//start数组的开始位置，k是剩余的个数，remain剩余的数
func dfs(start, k, remain int, path []int, res *[][]int) {

	if k <= 0 || k > 9-start || remain <= 0 {
		return
	}

	for i := start; i <= 9; i++ {
		if remain-i < 0 {
			break
		}
		path = append(path, i)
		if k-1 == 0 && remain-i == 0 {
			(*res) = append((*res), path)
			break
		}
		dfs(i+1, k-1, remain-i, path, res)
		path = path[:len(path)-1]
	}
}
