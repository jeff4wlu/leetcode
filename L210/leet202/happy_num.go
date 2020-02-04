package leet202

func HappyNum(num int) bool {

	dict := map[int]int{}

	for {
		var a, b, res int
		for num != 0 {
			a = num / 10
			b = num % 10
			res += b *b
			num = a
		}
		if res == 1 {
			return true
		}

		if _, ok := dict[res]; ok {
			return false
		}
		dict[res]++
		num = res
	}

}
