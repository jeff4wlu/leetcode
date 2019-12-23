package leet38

import "strconv"

func countAndSayDFS(s string) string {

	n := len(s)
	if n == 1 {
		return "1" + s[0:]
	}
	var res string
	cnt := 1
	for i := n - 1; i > 0; i-- {
		if s[i:i+1] == s[i-1:i] {
			cnt++
		} else {
			res = strconv.Itoa(cnt) + s[i:i+1] + res
			cnt = 1
		}
	}

	return strconv.Itoa(cnt) + s[0:1] + res
}

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := "1"
	for i := 1; i < n; i++ {
		res = countAndSayDFS(res)
	}

	return res
}
