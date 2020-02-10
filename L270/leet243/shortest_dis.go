package leet243

import "leetcode/infra"

func ShortestDis(dict []string, s, t string) int {

	n := len(dict)
	min := infra.INT32_MAX

	p1, p2 := -1, -1

	for i := 0; i < n; i++ {
		if dict[i] == s {
			p1 = i
		} else if dict[i] == t {
			p2 = i
		}
		if p1 != -1 && p2 != -1 {
			tmp := p1 - p2
			if tmp < 0 {
				tmp = -tmp
			}
			min = infra.IntMin(min, tmp)
		}
	}

	return min
}
