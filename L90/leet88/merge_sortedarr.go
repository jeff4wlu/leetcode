package leet88

func MergeSortedArr(a, b []int, m, n int) {

	pos := m + n - 1
	for i := n - 1; i >= 0; i-- { //b串

		for j := m - 1; j >= 0; j-- { //a串
			if b[i] >= a[j] {
				a[pos] = b[i]
				pos--
				break
			} else {
				a[pos] = a[j]
				m--
				pos--
			}
		}
	}
}
