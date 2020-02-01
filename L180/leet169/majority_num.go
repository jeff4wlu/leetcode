package leet169

//先排序，求中间那个元素
func MajorityNum1(in []int) int {

	n := len(in)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
	return in[(n-1)/2]
}

//使用bit操作。众数在每一位上的1肯定也是结果的1,所以res肯定就是众数
func MajorityNum2(in []int) int {

	n := len(in)
	var res int
	for i := 0; i < 32; i++ {
		var c0, c1 int
		for j := 0; j < n; j++ {
			if (in[j] >> uint8(i) & 1) == 1 {
				c1++
			} else {
				c0++
			}
		}
		if c1 > c0 {
			res += 1 << uint8(i)
		}
	}
	return res
}
