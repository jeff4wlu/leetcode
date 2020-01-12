package leet93

import "strconv"

//ip的规律是分四段，每一段可以有1、2或3个选择。且每段不能大于255且不能0开头
func RestoreIPAddr(s string) []string {

	if len(s) > 12 {
		return nil
	}

	res := []string{}
	var path string
	dfs(s, 4, path, &res)

	return res

}

//lvl还剩下多少段
func dfs(s string, lvl int, path string, res *[]string) {

	n := len(s)

	if lvl == 0 {
		if n == 0 {
			path = path[:len(path)-1]
			*res = append(*res, path)
		}
		return
	}

	least := n / lvl
	if least == 0 {
		return
	}

	if s[0] != '0' {
		tmp := path
		//选择一位
		path += s[:1]
		path += "."
		dfs(s[1:], lvl-1, path, res)
		path = tmp

		//选择两位
		if n >= 2 {
			path += s[:2]
			path += "."
			dfs(s[2:], lvl-1, path, res)
			path = tmp
		}
		//选择三位
		if n >= 3 {
			num, _ := strconv.Atoi(s[:3])
			if num <= 255 {
				path += s[:3]
				path += "."
				dfs(s[3:], lvl-1, path, res)
			}

		}

	}

}
