package leet191

func NumsBit(num uint32) int {

	var res int
	for i := 0; i < 32; i++ {
		if num>>uint(i)&1 == 1 {
			res++
		}
	}

	return res
}
