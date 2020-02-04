package leet190

import "fmt"

func ReverseBit(num uint32) uint32 {

	var res uint32
	for i := 0; i < 32; i++ {
		//先求出num中每一位是0还是1，然后再左移到对应位置
		res |= (num >> uint(i) & 1) << (31 - uint(i))
		fmt.Printf("--%032b\n", res)
	}
	return res
}
