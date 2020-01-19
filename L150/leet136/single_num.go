package leet136

func SingleNum(in []int) int {

	res := in[0]

	for i := 1; i < len(in); i++ {
		res ^= in[i]
	}

	return res
}
