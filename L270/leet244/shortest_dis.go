package leet244

import "leetcode/infra"

func ShortestDis(words []string, s, t string) int {

	//n := len(dict)
	dict := map[string][]int{}
	for i, v := range words {
		dict[v] = append(dict[v], i)
	}

	sarr := dict[s]
	tarr := dict[t]

	var i, j int
	res := infra.INT32_MAX
	//O(mn)优化为O(n+m), aabb,abab模式，求出ab之间的最短距离，只需要3次，非4次
	for i < len(sarr) && j < len(tarr) {
		tmp := sarr[i] - tarr[j]
		if tmp < 0 {
			tmp = -tmp
			i++
		} else {
			j++
		}
		res = infra.IntMin(tmp, res)
	}
	return res
}
