package leet34

func FirstLastSortedArr(in []int, target int) []int {

	n := len(in)

	var left, right, start, end int
	right = n - 1

	for left <= right {
		mid := left + (right-left)/2
		if in[mid] == target {
			for i := mid - 1; i >= left; i-- {
				if in[i] != target {
					start = i + 1
					break
				}
			}

			for i := mid + 1; i <= right; i++ {
				if in[i] != target {
					end = i - 1
					break
				}
			}

			return []int{start, end}

		} else if in[mid] < target {
			left = mid + 1
		} else {

			right = mid - 1
		}
	}

	return []int{-1, -1}
}
