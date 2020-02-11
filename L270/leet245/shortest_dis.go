package leet245

import "leetcode/infra"

func ShortestDis(dict []string, s, t string) int {

	n := len(dict)
	min := infra.INT32_MAX

	p1, p2 := -1, -1

	if s != t {
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
	} else {
		p1 = infra.INT32_MIN //必须确保有距离，如果只有一个点，本算法出错
		for i := 0; i < n; i++ {
			if dict[i] == s {
				min = infra.IntMin(min, i-p1)
				p1 = i
			}
		}
	}

	return min
}
