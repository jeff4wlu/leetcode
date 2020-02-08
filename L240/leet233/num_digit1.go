package leet233

func NumDigitOne(num int) int {

	if num <= 0 {
		return 0
	}

	var cnt int

	for i := 1; i <= num; i++ {
		quotient, remainder := i, -1
		for quotient != 0 {
			remainder = quotient % 10 //求出个位上的数
			quotient = quotient / 10  //求去掉各位后的数
			if remainder == 1 {
				cnt++
			}
		}
	}

	return cnt
}
