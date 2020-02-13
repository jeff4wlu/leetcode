package leet259

func SumSmaller(nums []int, target int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	path := []int{}
	res := [][]int{}
	dfs(nums, path, 0, target, &res)
	return res
}

func dfs(nums []int, path []int, pos, target int, res *[][]int) {

	n := len(path)
	if n == 3 {
		tmp := make([]int, 3)
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}

	for i := pos; i < len(nums); i++ {
		tmp := target - nums[i]
		if tmp > 0 {
			path = append(path, nums[i])
			dfs(nums, path, i+1, tmp, res)
			path = path[:len(path)-1]
		}

	}
}
