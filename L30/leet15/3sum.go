package leet15

import "sort"

func ThreeSum(in []int) [][]int {

	var res [][]int

	if len(in) < 3 {
		return nil
	}
	sort.Ints(in)

	if in[0] > 0 {
		return nil
	}

	for k := 0; k < len(in)-2; k++ {
		if in[k] > 0 {
			break
		}
		if k > 0 && in[k] == in[k-1] {
			continue
		}
		target := 0 - in[k]
		i := k + 1
		j := len(in) - 1
		for i < j {
			if target == in[i]+in[j] {
				tmp := []int{in[k], in[i], in[j]}
				res = append(res, tmp)
				for i < j && in[i] == in[i+1] {
					i++
				}
				for i < j && in[j] == in[j-1] {
					j--
				}
				i++
				j--

			} else if target > in[i]+in[j] {
				i++
			} else {
				j--
			}

		}

	}

	return res
}
