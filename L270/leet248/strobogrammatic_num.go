package leet248

import "strconv"

func StrobogrammaticNum(p, q string) int {

	pl := len(p)
	ql := len(q)
	pv, _ := strconv.Atoi(p)
	qv, _ := strconv.Atoi(q)

	var cpy []string
	re := []string{}
	n1 := []string{"0", "1", "8"}
	n2 := []string{"11", "69", "88", "96"}
	op := []string{"11", "69", "88", "96", "00"}

	cpy = make([]string, 4)
	copy(cpy, n2)
	for j := 3; j <= ql; j++ { //数字长度
		a := j / 2
		b := j % 2

		if b == 0 {
			for i := 1; i < a; i++ {
				tmp := []string{}
				for _, v1 := range cpy {
					for _, v2 := range op {
						s := v1[:i] + v2 + v1[i:]
						tint, _ := strconv.Atoi(s)
						if tint >= pv && tint < qv || tint == qv {
							re = append(re, s)
						}
						tmp = append(tmp, s)
					}
				}
				cpy = tmp
			}

		} else {
			for i := 1; i <= a; i++ {
				tmp := []string{}
				for _, v1 := range cpy {
					for _, v2 := range n1 {
						s := v1[:i] + v2 + v1[i:]
						tint, _ := strconv.Atoi(s)
						if tint >= pv && tint < qv || tint == qv {
							re = append(re, s)
						}
						tmp = append(tmp, s)
					}
				}
				cpy = tmp
			}
		}
	}
	if pl < 3 {
		for _, v := range n1 {
			tint, _ := strconv.Atoi(v)
			if tint >= pv && tint < qv || tint == qv {
				re = append(re, v)
			}
		}
		for _, v := range n2 {
			tint, _ := strconv.Atoi(v)
			if tint >= pv && tint < qv || tint == qv {
				re = append(re, v)
			}
		}
	}

	return len(re)
}
