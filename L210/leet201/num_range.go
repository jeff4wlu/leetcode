package leet201

//求m,n中最左边相同的1
func NumRange(m, n int) int {

	tmp := []int{}
	for i := 0; i < 32; i++ {
		a := m >> uint(i) & 1
		b := n >> uint(i) & 1
		if a == b && a == 1 {
			tmp = append(tmp, i)
		}
	}

	length := len(tmp)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1 << uint(tmp[length])
	}

	start := tmp[length-1]
	var end int
	for i := length - 2; i >= 0; i-- {
		if start != tmp[i]+1 {
			end = tmp[i+1]
			break
		}
	}

	var res int
	for i := start; i >= end; i-- {
		res |= 1 << uint(i)
	}
	return res
}
