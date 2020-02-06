package leet214

import "leetcode/infra"

func ShortestPalindrome(s string) string {

	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}
	if n == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[1]) + s
		}
	}

	var left, right int
	res := []string{}
	for i := 1; i < n-1; i++ {
		left, right = i, i
		isPalindrome := true
		for left >= 0 && right < n {
			if s[left] != s[right] {
				isPalindrome = false
				break
			}
			left--
			right++
		}

		if isPalindrome {
			if left < 0 && right >= n { //本字符串就是回文
				return s
			}
			if left < 0 {
				gap := n - right
				tmp := s
				for gap > 0 {
					tmp = string(s[right]) + tmp
					gap--
					right++
				}
				res = append(res, tmp)
			} else {
				gap := left + 1
				tmp := s
				for gap > 0 {
					tmp = tmp + string(s[left])
					gap--
					left--
				}
				res = append(res, tmp)
			}
		}
	}

	min := infra.INT32_MAX
	var re string
	for _, v := range res {
		if len(v) < min {
			min = len(v)
			re = v
		}
	}
	if re != "" {
		return re
	}
	re = s
	for i := 1; i < n; i++ {
		re = string(s[i]) + re
	}
	return re
}
