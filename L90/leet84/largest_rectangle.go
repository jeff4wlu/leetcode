package leet84

//n*n的时间复杂度，遍历
func LargestRect(in []int) int {

	var max int
	n := len(in)
	for i := n - 1; i >= 0; i-- {
		tmp := interateCurHeight(i, in)
		if max < tmp {
			max = tmp
		}
	}
	return max
}

//对算法1进行剪枝，只计算峰值的遍历。峰值就是当下一个比当前小时，当前是峰值。但时间复杂度仍然是n*n
func LargestRect1(in []int) int {

	var max int
	n := len(in)

	for i := 0; i < n-1; i++ {
		if in[i] > in[i+1] {
			tmp := interateCurHeight(i, in)
			if max < tmp {
				max = tmp
			}

		}
	}

	if in[n-1] > in[n-2] {
		tmp := interateCurHeight(n-1, in)
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func interateCurHeight(idx int, in []int) int {
	var max int
	shortest := in[idx]
	width := 2
	if max < in[idx] {
		max = in[idx]
	}
	for j := idx - 1; j >= 0; j-- {

		if in[j] < shortest {
			shortest = in[j]
		}
		if max < shortest*width {
			max = shortest * width
		}
		width++
	}
	return max
}
