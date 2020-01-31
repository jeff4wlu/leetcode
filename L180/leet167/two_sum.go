package leet167

//O(nlg(n))
func TwoSum(in []int, target int) []int {
	n := len(in)
	if n < 2 {
		return []int{}
	}
	if n == 2 {
		if in[0]+in[1] != target {
			return []int{}
		}
		return []int{1, 2}
	}

	for i := 0; i < n-1; i++ {
		tmp := target - in[i]
		a := two(in, i+1, n-1, tmp)
		if a != -1 {
			return []int{i + 1, a + 1}
		}
	}

	return []int{}
}

func two(in []int, left, right, target int) int {
	n := len(in)
	if n < 1 {
		return -1
	}
	if n == 1 {
		return left
	}

	mid := left + (right-left)/2
	for left < right {
		if in[mid] == target {
			return mid
		}
		if in[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right

}

func TwoSum2(in []int, target int) []int {
	n := len(in)
	if n < 2 {
		return []int{}
	}
	if n == 2 {
		if in[0]+in[1] != target {
			return []int{}
		}
		return []int{1, 2}
	}

	left, right := 0, n-1
	for left < right {
		if in[left]+in[right] == target {
			return []int{left + 1, right + 1}
		}
		if in[left]+in[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
