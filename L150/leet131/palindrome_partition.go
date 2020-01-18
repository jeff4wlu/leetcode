package leet131

import "leetcode/L150/leet125"

func PalindromePartiion(s string) [][]string {
	path := []string{}
	res := [][]string{}
	dfs(s, path, &res)
	return res
}

func dfs(s string, path []string, res *[][]string) {
	n := len(s)

	if n <= 1 {
		if n == 1 {
			path = append(path, s)
		}
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)

		return
	}

	for i := 1; i < n; i++ {
		isPalindrome := leet125.ValidPalindrome(s[:i])
		if isPalindrome {
			path = append(path, s[:i])
			dfs(s[i:], path, res)
			path = path[:len(path)-1]
		}
	}
}
