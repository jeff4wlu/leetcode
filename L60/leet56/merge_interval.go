package leet56

func MergeInterval(in [][]int) [][]int {

	size := len(in)

	start := in[0][0]
	end := in[0][1]

	res := [][]int{}

	for i := 1; i < size; i++ {
		//全包含在内
		if in[i][1] <= end {
			continue
		}
		//全不包含
		if in[i][0] > end {
			res = append(res, []int{start, end})
			start = in[i][0]
			end = in[i][1]
			continue
		}

		//有重叠的地方
		end = in[i][1]

	}
	res = append(res, []int{start, end})
	return res
}
