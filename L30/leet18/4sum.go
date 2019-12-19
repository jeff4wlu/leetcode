package leet18

import "sort"

func FourSum(in []int, target int) [][]int {

	nums := len(in)

	if nums < 4 {
		return nil
	}

	sort.Ints(in)

	var res [][]int

	for i := 0; i < nums-3; i++ {
		if i > 0 && in[i] == in[i-1] {
			continue
		}
		for j := i + 1; j < nums-2; j++ {
			if j > i+1 && in[j] == in[j-1] {
				continue
			}
			left := j + 1
			right := nums - 1
			for left < right {
				sum := in[i] + in[j] + in[left] + in[right]
				if sum == target {
					res = append(res, []int{in[i], in[j], in[left], in[right]})
					for left < right && in[left] == in[left+1] {
						left++
					}
					for left < right && in[right] == in[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}

	}

	return res

}
