package leet139

import "leetcode/infra"

func WordBreak(s string, dictArr []string) bool {

	dict := make(map[string]int)

	var max int
	min := infra.INT32_MAX
	for _, v := range dictArr {
		dict[v]++
		if min > len(v) {
			min = len(v)
		}
		if max < len(v) {
			max = len(v)
		}
	}

	path := []string{}
	res := [][]string{}
	dfs(s, min, max, dict, path, &res)
	if len(res) > 0 {
		return true
	}

	return false
}

func dfs(s string, min, max int, dict map[string]int, path []string, res *[][]string) {

	if s == "" {
		return
	}

	if _, ok := dict[s]; ok {
		path = append(path, s)
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := min; i <= max && i <= len(s); i++ {
		if _, ok := dict[s[:i]]; ok {
			path = append(path, s[:i])
			dfs(s[i:], min, max, dict, path, res)
			path = path[:len(path)-1]
		}
	}
}
