package leet78

//dfs,有点像暴力破解，回溯。题目要求按升序排。
//原集合中每一个元素在子集中有两种状态：要么存在、要么不存在。这样构造子集的过程中每个元素就有两种选择方法：选择、不选择，
//因此可以构造一颗二叉树，例如对于例子中给的集合[1,2,3]，构造的二叉树如下（左子树表示选择该层处理的元素，右子树不选择），最后得到的叶子节点就是子集：
func Subsets(sets []int) [][]int {

	res := [][]int{}
	path := []int{}

	dfs(sets, 0, path, &res)

	return res
}

func dfs(in []int, idx int, path []int, res *[][]int) {

	if idx == len(in) { //已经没有元素来做决定了

		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	//不选择此元素
	dfs(in, idx+1, path, res)

	//选择此元素做排列
	path = append(path, in[idx])
	dfs(in, idx+1, path, res)

}
