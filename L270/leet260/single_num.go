package leet260

func SingleNum(nums []int) []int {

	res := []int{}
	var tmp int
	for _, v := range nums {
		tmp ^= v
	}

	var mask int
	for i := 0; i < 32; i++ {
		mask = tmp & (1 << uint(i))
		if mask != 0 {
			break
		}
	}

	ones := []int{}
	twos := []int{}
	for _, v := range nums {
		if v&mask == 0 {
			twos = append(twos, v)
		} else {
			ones = append(ones, v)
		}
	}

	for _, v := range ones {
		tmp ^= v
	}
	res = append(res, tmp)
	tmp = 0
	for _, v := range twos {
		tmp ^= v
	}
	res = append(res, tmp)
	return res
}
