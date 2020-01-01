package leet66

func PlusOne(num []int) []int {
	n := len(num)

	if n < 1 {
		return nil
	}

	var adv, total int
	res := []int{}

	total = num[n-1] + 1
	if total >= 10 {
		adv = 1
		res = append(res, total-10)
	} else {
		res = append(res, total)
	}
	i := n - 2

	for ; i >= 0; i-- {

		total = num[i] + adv
		if total >= 10 {
			adv = 1
			tmp := []int{total - 10}
			res = append(tmp, res...)
		} else {
			tmp := []int{total}
			res = append(tmp, res...)
			break
		}

	}
	i--

	if i >= 0 {
		res = append(num[:i+1], res...)
	} else if i < 0 && adv == 1 {
		res = append([]int{1}, res...)
	}
	return res
}
