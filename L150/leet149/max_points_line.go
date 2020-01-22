package leet149

func MaxPointsLine(in [][]int) int {

	n := len(in)
	if n <= 1 {
		return n
	}

	var max int

	for _, v := range in {
		dict := map[int]int{}
		xdict := map[int]int{}
		ydict := map[int]int{}
		for _, k := range in {
			if v[0] != k[0] && v[1] != k[1] {
				x := v[0] - k[0]
				y := v[1] - k[1]
				gcd := maxCommonDividor(x, y)
				dx := x / gcd
				dy := y / gcd
				key := hashCode(dx, dy)
				dict[key]++
			} else if v[0] == k[0] && v[1] == k[1] {
				continue
			} else if v[0] == k[0] {
				xdict[v[0]]++
			} else {
				ydict[v[1]]++
			}

		}
		for _, j := range dict {
			if max < j {
				max = j
			}
		}
		for _, j := range xdict {
			if max < j {
				max = j
			}
		}
		for _, j := range ydict {
			if max < j {
				max = j
			}
		}
	}

	return max + 1
}

//被公约数后的分子分母做散列出一个key，方便查找
func hashCode(dx, dy int) int {
	return (dx << 32) ^ dy
}

//求最大公约数，避免整数除法精度问题导致hashmap找不到相同的key
//m，n中不能有0
func maxCommonDividor(m, n int) int {
	if m < n {
		m, n = n, m
	}

	for m%n != 0 {
		tmp := m % n
		m = n
		n = tmp
	}

	return n
}
