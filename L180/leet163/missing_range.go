package leet163

import "strconv"

func MissingRange(in []int, lower, upper int) []string {

	n := len(in)
	if n == 0 {
		return []string{}
	}

	l := lower
	res := []string{}
	for i := 0; i < n; i++ {
		tmp := in[i] - l
		if i == 0 && tmp == 1 {
			res = append(res, srange(0, 0))
		}
		if tmp > 1 {
			res = append(res, srange(l+1, in[i]-1))
		}

		l = in[i]
	}
	tmp := upper - in[n-1]
	if tmp != 0 {
		res = append(res, srange(in[n-1]+1, 99))
	}

	return res

}

func srange(lower, upper int) string {
	if lower == upper {
		return strconv.Itoa(lower)
	}

	return strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
}
