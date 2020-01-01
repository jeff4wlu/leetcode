package leet66

//因为只加一，所以可以利用9来判断进位
func PlusOne(num []int) []int {

	n := len(num)

	for i := n - 1; i >= 0; i-- {
		if num[i] == 9 {
			num[i] = 0
		} else {
			num[i]++
			break
		}
	}

	res := []int{}
	if num[0] == 0 {
		res = append([]int{1}, num...)
	} else {
		res = append(res, num...)
	}

	return res
}
