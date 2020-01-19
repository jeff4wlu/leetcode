package leet135

func Candy(rate []int) int {

	candle := make([]int, len(rate))

	var count int
	for i := 0; i < len(rate); i++ {
		count++
		candle[i]++
		if i > 0 {
			if rate[i] > rate[i-1] {
				count++
				candle[i] = candle[i-1] + 1
			}
		}
	}

	for i := len(rate) - 1; i >= 0; i-- {
		if i < len(rate)-2 {
			if rate[i] > rate[i+1] {
				count++
				candle[i] = candle[i+1] + 1
			}
		}
	}

	return count
}
