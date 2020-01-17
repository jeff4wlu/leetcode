package leet125

import "strings"

//双指针
func ValidPalindrome(s string) bool {

	n := len(s)
	if n <= 1 {
		return true
	}

	s = strings.ToLower(s)

	j := n - 1
	i := 0
	for i <= j {

		for s[i] < 'a' || s[i] > 'z' {
			i++
			continue
		}
		for s[j] < 'a' || s[j] > 'z' {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
