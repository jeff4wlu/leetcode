package leet12

var roman = [7]byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
var value = [7]int{1000, 500, 100, 50, 10, 5, 1}

func IntToRoman(num int) string {

	//res := make([]byte, 20)
	var res []byte

	for i := 0; i < 7; i += 2 {
		//获取各个位置上的数字
		x := num / value[i]

		switch {
		case x < 4:
			for j := 1; j <= x; j++ {
				res = append(res, roman[i])
			}
		case x == 4:
			res = append(res, roman[i])
			res = append(res, roman[i-1])
		case x > 4 && x < 9:
			res = append(res, roman[i-1])
			for j := 6; j <= x; j++ {
				res = append(res, roman[i])
			}
		case x == 9:
			res = append(res, roman[i])
			res = append(res, roman[i-2])
		}

		num %= value[i]

	}
	return string(res)
}
