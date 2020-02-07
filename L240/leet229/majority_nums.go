package leet229

//使用摩尔投票法，nums内不能有0，因为0是默认值，会干扰结果
func MajorityNums(nums []int) []int {

	res := []int{}
	n := len(nums)

	if n < 3 {
		return res
	}

	var a, b, cnta, cntb int

	for i := 0; i < n; i++ {
		if a == nums[i] {
			cnta++
		} else if b == nums[i] {
			cntb++
		} else if cnta == 0 {
			cnta = 1
			a = nums[i]
		} else if cntb == 0 {
			cntb = 1
			b = nums[i]
		} else {
			cnta--
			cntb--
		}
	}

	cnta, cntb = 0, 0
	for i := 0; i < n; i++ {
		if a == nums[i] {
			cnta++
		}
		if b == nums[i] {
			cntb++
		}
	}

	if cnta > int(n/3) {
		res = append(res, a)
	}
	if cntb > int(n/3) {
		res = append(res, b)
	}

	return res

}
