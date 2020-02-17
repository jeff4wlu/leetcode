package leet271

import (
	"leetcode/infra"
	"strconv"
)

func Encode(words []string) string {

	var res string
	for _, v := range words {
		res += encode(v)
	}

	return res
}

func Decode(s string) []string {
	n := len(s)
	if n <= 3 {
		return []string{}
	}

	res := []string{}

	var start int
	for i := 0; i < n; {
		for ; s[i] >= '0' && s[i] <= '9'; i++ {
			continue
		}
		if start == i {
			return []string{}
		}
		if s[i] != '/' {
			return []string{}
		}
		length, err := strconv.Atoi(s[start:i])
		if err != nil {
			return []string{}
		}

		res = append(res, s[i+1:i+length+1])
		start = i + length + 1
		i = start

	}

	return res
}

func encode(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	if n >= infra.INT32_MAX-11 {
		return ""
	}

	return strconv.Itoa(n) + "/" + s
}
