package leet75

// 借鉴三路快排中的划分思路
func SortCols(in *[]int) {

	i, j, k := 0, 0, len(*in)-1

	// for 循环中， nums[i:j] 中始终全是 1
	// 循环结束后，
	// nums[:i] 中全是 0
	// nums[j:] 中全是 2
	for j <= k {
		switch (*in)[j] {
		case 0:
			(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
			i++
			j++
		case 1:
			j++
		case 2:
			(*in)[j], (*in)[k] = (*in)[k], (*in)[j]
			k--
		}
	}
}
