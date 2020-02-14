package leet259

import "leetcode/infra"

func SumSmaller1(nums []int, target int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	res := [][]int{}

	//从小到大
	nums = infra.BubbleSort(nums, false)

	for i := 0; i <= n-3; i++ {
		small, big := i+1, n-1
		for small < big {
			if nums[i]+nums[small]+nums[big] < target {
				for small < big {
					res = append(res, []int{nums[i], nums[small], nums[big]})
					big--
				}
				break
			}
			big--
		}
	}

	return res
}

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
