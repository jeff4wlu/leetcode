package leet157

import "leetcode/infra"

func ReadN(f []byte, n int) int {

	if n <= 4 {
		return read4(f)
	}

	var res int
	for n-res > 0 {
		tmp := read4(f[res:])
		if tmp == 0 {
			break
		}
		res += tmp
	}

	return infra.IntMin(n, res)
}

func read4(f []byte) int {
	n := len(f)
	if n < 4 {
		return n
	}

	return 4
}
