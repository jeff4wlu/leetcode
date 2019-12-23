package leet39

//找所有可能解，需使用递归回溯的算法。
//leaf是解空间的边界，此题为遍历完成数组中所有元素
//good leaf node，各个加数之和等于target
func CombinationSum(nums []int, target int) [][]int {

	res := [][]int{}
	path := []int{} //这里分配了内存，所以不会有nil的情况出现
	CombinationSumDSF(nums, path, target, &res)
	return res
}

//nums,供计算的数组
//target,当前的和
//path,当前的各个加数
//res,结果集
func CombinationSumDSF(nums, path []int, target int, res *[][]int) {

	if target == 0 {
		*res = append(*res, path)
		return
	}

	if len(nums) == 0 || target < nums[0] {
		return
	}

	path = path[:len(path):len(path)]

	//相同元素相加
	CombinationSumDSF(nums, append(path, nums[0]), target-nums[0], res)

	//步进
	CombinationSumDSF(nums[1:], path, target, res)

}
