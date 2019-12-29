package leet57

func InsertInterval(in [][]int, a []int) [][]int {

	res := [][]int{}
	pre := a
	for i := 0; i < len(in); i++ {
		if pre[1] < in[i][0] { //插入点左边
			res = append(res, pre)
			res = append(res, in[i:]...)
			return res
		} else if pre[0] > in[i][1] { //插入点右边
			res = append(res, in[i])

		} else { //重叠
			if pre[0] > in[i][0] {
				pre[0] = in[i][0]
			}
			if pre[1] < in[i][1] {
				pre[1] = in[i][1]
			}
		}
	}

	return append(res, pre)

}
