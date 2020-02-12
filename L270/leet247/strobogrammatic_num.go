package leet247

var dict []byte = []byte{'1', '6', '8', '9', '6', '8', '9', '1'}

func StrobogrammaticNum(k int) []string {

	var res []string
	n1 := []string{"0", "1", "8"}
	n2 := []string{"11", "69", "88", "96"}
	op := []string{"11", "69", "88", "96", "00"}
	if k == 0 {
		return []string{}
	}
	if k == 1 {
		return n1
	}
	if k == 2 {
		return n2
	}

	a := k / 2
	b := k % 2
	res = make([]string, 4)
	copy(res, n2)
	if b == 0 {

		for i := 1; i < a; i++ {
			tmp := []string{}
			for _, v1 := range res {
				for _, v2 := range op {
					s := v1[:i] + v2 + v1[i:]
					tmp = append(tmp, s)
				}
			}
			res = tmp
		}

	} else {
		for i := 1; i <= a; i++ {
			tmp := []string{}
			for _, v1 := range res {
				for _, v2 := range n1 {
					s := v1[:i] + v2 + v1[i:]
					tmp = append(tmp, s)
				}
			}
			res = tmp
		}
	}

	return res
}
