package leet31

func NextPermutation(in []int) {

	n := len(in)

	isMax := true
	for i := n - 1; i > 0; i-- {
		if in[i-1] < in[i] {
			for j := n - 1; j > i-1; j-- {
				if in[j] > in[i-1] {
					in[j], in[i-1] = in[i-1], in[j]
					reverseSortedInt(in[i:]) //slice是相当于穿引用
					return
				}
			}

		}

	}

	if isMax {
		reverseSortedInt(in)
		return
	}

}

func reverseSortedInt(in []int) {
	n := len(in)
	mid := n / 2
	for i := 0; i < mid; i++ {
		in[i], in[n-1-i] = in[n-1-i], in[i]
	}
}
