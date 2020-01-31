package leet166

import "strconv"

func FractionDec(num, denominator int) string {

	var res, tmp string
	if num*denominator < 0 {
		res += "-"
	}
	if num < 0 {
		num = -num
	}
	if denominator < 0 {
		denominator = -denominator
	}
	a := num / denominator
	b := num % denominator
	if b == 0 {
		return strconv.Itoa(a)
	}

	res += strconv.Itoa(a)

	dict := map[int]int{}
	recur := false
	for b != 0 {
		a = b * 10 / denominator
		b = b * 10 % denominator
		dict[a]++
		if dict[a] > 1 {
			recur = true
			break
		}
		tmp += strconv.Itoa(a)
	}

	c := strconv.Itoa(a)

	i := len(tmp) - 1
	for ; i >= 0; i-- {
		if tmp[i] != c[0] {
			break
		}
	}

	res += "."
	res += tmp[:i+1]
	if recur {
		res += "("
	}

	res += tmp[i+1:]
	if recur {
		res += ")"
	}

	return res
}
