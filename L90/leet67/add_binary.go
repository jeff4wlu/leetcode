package leet67

import "strconv"

func AddBinary(a, b string) string {

	if len(a) < len(b) {
		a, b = b, a
	}

	res := add(toIntArr(a, len(a)), toIntArr(b, len(a)))

	return intsToStr(res)

}

func intsToStr(in []int) string {
	var res string
	for _, v := range in {
		res += strconv.Itoa(v)
	}
	return res
}

func toIntArr(s string, l int) []int {
	res := make([]int, l)

	for i, v := range s {
		res[l-len(s)+i] = int(v - '0')
	}
	return res
}

func add(a, b []int) []int {
	res := make([]int, len(a)+1)
	carried := 0
	for i := len(a) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + carried
		res[i+1] = sum % 2
		carried = sum / 2
	}

	res[0] = carried

	if res[0] == 0 {
		return res[1:]
	}
	return res
}
