package leet122

//卖不能在买前发生，且买卖一次完成后才能发生第二次
func BestTimeStock2(in []int) int {

	if len(in) < 2 {
		return 0
	}

	var res int

	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			res += in[i] - in[i-1]
		}
	}

	return res
}
