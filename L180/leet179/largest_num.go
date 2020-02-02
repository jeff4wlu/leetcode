package leet179

import "strconv"

func LargestNum(in []int) string {

	s := []string{}
	for _, v := range in {
		tmp := strconv.Itoa(v)
		s = append(s, tmp)
	}

	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j > i; j-- {
			if !greater(s[j-1], s[j]) {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}

	var res string
	for _, v := range s {
		res += v
	}
	return res
}

func greater(a, b string) bool {

	as := a + b
	bs := b + a
	for i := 0; i < len(as); i++ {
		tmp := int(as[i]) - int(bs[i])
		if tmp > 0 {
			return true
		}
		if tmp < 0 {
			return false
		}
	}
	return true
}
