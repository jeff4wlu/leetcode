package leet91

import "strconv"

//dfs穷举和回溯法。也可以用dp
//任何算法只要可以使用两种选择来穷举的，都可以用二叉树的dfs。参见leet89,90
func DecodeWays(s string) int {

	res := []string{}
	var path string

	dfs(s, 0, path, &res)

	return len(res)
}

func dfs(s string, idx int, path string, res *[]string) {

	n := len(s)

	if idx == n {
		tmp := path
		*res = append(*res, tmp)
		return
	}

	if s[idx] == '0' {
		dfs(s, idx+1, path, res)
		return
	}

	//选择一位数来decode
	tmp := path
	path += string(s[idx])
	dfs(s, idx+1, path, res)
	path = tmp

	//选择两位数来decode

	if idx+1 < n {
		num, _ := strconv.Atoi(s[idx : idx+2])
		if num <= 26 {
			path += s[idx : idx+2]
			dfs(s, idx+2, path, res)
		}

	}
}
