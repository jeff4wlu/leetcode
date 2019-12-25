package leet46

//分治模板
func PermutationsDSF(nums, path []int, res *[][]int) {

	n := len(nums)
	if n == 1 {
		path = append(path, nums[0])
		*res = append(*res, path)
		return
	}

	for i := 0; i < n; i++ {
		var tmp []int

		if i > 0 && i < n-1 {
			tmp = append(tmp, nums[0:i]...)
			tmp = append(tmp, nums[i+1:]...)
		} else if i == n-1 {
			tmp = append(tmp, nums[0:n-1]...)
		} else {
			tmp = append(tmp, nums[i+1:]...)
		}

		//tmp = tmp[:len(tmp):len(tmp)]
		PermutationsDSF(tmp, append(path, nums[i]), res)
	}
}

func Permutations(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
	PermutationsDSF(nums, path, &res)
	return res
}
