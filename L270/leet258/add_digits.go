package leet258

func AddDigits(num int) int {

	for num/10 > 0 {
		var sum int
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}
