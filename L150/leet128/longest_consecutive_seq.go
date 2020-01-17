package leet128

func LongestConsecutiveSeq(in []int) int {

	n := len(in)

	if n <= 1 {
		return n
	}

	dict := map[int]int{}

	for i := 0; i < n; i++ {
		dict[in[i]] = 0
	}

	var max int
	for i := 0; i < n; i++ {
		if dict[in[i]] > 0 {
			continue
		}
		if _, ok := dict[in[i]]; ok {
			count := 1
			tmp := in[i]
			for {
				tmp++
				if _, ok := dict[tmp]; ok {
					count++
					dict[tmp]++
					continue
				}
				break
			}
			tmp = in[i]
			for {
				tmp--
				if _, ok := dict[tmp]; ok {
					count++
					dict[tmp]++
					continue
				}
				break
			}
			if count > max {
				max = count
			}
		}

	}

	return max

}
