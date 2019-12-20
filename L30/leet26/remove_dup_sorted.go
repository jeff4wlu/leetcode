package leet26

func RemoveDupSortedArr(in []int) int {

	if len(in) < 2 {
		return len(in)
	}
	var a int

	for b := a + 1; b < len(in); b++ {
		if in[a] != in[b] {
			a++
			in[a], in[b] = in[b], in[a]
		}
	}

	return a + 1

}
