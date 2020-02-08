package leet238

func ProductArr(nums []int) []int {

	res := []int{}

	n := len(nums)
	if n < 2 {
		return res
	}

	ins := make([]int, n)
	des := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			ins[i] = 1
			des[n-i-1] = 1
			continue
		}
		ins[i] = nums[i-1] * ins[i-1]
		des[n-i-1] = nums[n-i] * des[n-i]
	}

	for i := 0; i < n; i++ {
		res = append(res, ins[i]*des[i])
	}

	return res
}
