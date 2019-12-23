package leet40

import "sort"

func CombinationSumUniq(nums []int, target int) [][]int {
	res := [][]int{}
	path := []int{} //这里分配了内存，所以不会有nil的情况出现
	sort.Ints(nums)
	combinationSumUniqDSF(nums, path, target, &res)
	return res
}

func combinationSumUniqDSF(nums, path []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, path)
		return
	}

	if len(nums) == 0 || target < nums[0] {
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	path = path[:len(path):len(path)]

	//本元素和后面的sum
	combinationSumUniqDSF(nums[1:], append(path, nums[0]), target-nums[0], res)

	//其他元素的sum值，注意跳开可能的相同元素，因为前面一句已经做好了
	combinationSumUniqDSF(next(nums), path, target, res)
}

func next(nums []int) []int {
	i := 0
	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}
	return nums[i+1:]
}
