package leet90

func Subsets(sets []int) [][]int {

	res := [][]int{}
	path := []int{}

	dfs(sets, 0, path, &res)

	return res
}

//此处要花二叉树图来找规律
func dfs(in []int, idx int, path []int, res *[][]int) {

	if idx == len(in) { //已经没有元素来做决定了

		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	firstSame := idx
	for firstSame >= 0 && in[firstSame] == in[idx] {
		firstSame--
	}
	firstSame++
	sameNum := idx - firstSame

	//不选择此元素
	dfs(in, idx+1, path, res)

	if sameNum == 0 || len(path) >= sameNum && path[len(path)-sameNum] == in[idx] {
		//选择此元素做排列
		path = append(path, in[idx])
		dfs(in, idx+1, path, res)
	}

}
