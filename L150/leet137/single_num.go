package leet137

func SingleNum(in []int) int {

	var a, b int

	for _, v := range in {
		tmp := a&(^b)&(^v) | (^a)&b&v
		b = (^a)&b&(^v) | (^a)&(^b)&v
		a = tmp
	}
	return b
}
