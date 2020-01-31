package leet166

import (
	"fmt"
	"strconv"
)

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
	for b != 0 {
		a = b * 10 / denominator
		b = b * 10 % denominator
		if _, ok := dict[a]; ok {
			pos := dict[a]
			return fmt.Sprintf("%s.%s(%s)", res, string(tmp[:pos]), string(tmp[pos:]))
		}

		tmp += strconv.Itoa(a)
		dict[a] = len(tmp) - 1
	}

	return fmt.Sprintf("%s.%s", res, string(tmp))
}
