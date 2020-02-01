package leet172

func FactorialTailing(num int) int {

	if num <= 4 {
		return 0
	}

	var res int
	for i := 5; i <= num; i++ {
		tmp := i
		var a int
		for tmp >= 5 {
			a = tmp / 5
			b := tmp % 5
			if b != 0 {
				break
			}
			res++
			tmp=a
		}
	}
	return res
}
