package leet273

var dict map[int]string = map[int]string{}

func init() {
	dict[0] = ""
	dict[1] = "One"
	dict[2] = "Two"
	dict[3] = "Three"
	dict[4] = "Four"
	dict[5] = "Five"
	dict[6] = "Six"
	dict[7] = "Seven"
	dict[8] = "Eight"
	dict[9] = "Nine"
	dict[10] = "Ten"

	dict[11] = "Eleven"
	dict[12] = "Twelve"
	dict[13] = "Thirteen"
	dict[14] = "Fourteen"
	dict[15] = "Fifteen"
	dict[16] = "Sixteen"
	dict[17] = "Seventeen"
	dict[18] = "Eighteen"
	dict[19] = "Nineteen"

	dict[20] = "Twenty"
	dict[30] = "Thirty"
	dict[40] = "Forty"
	dict[50] = "Fifty"
	dict[60] = "Sixty"
	dict[70] = "Seventy"
	dict[80] = "Eighty"
	dict[90] = "Ninety"

}

func Int2Eng(num int) string {

	var res string

	if num == 0 {
		return "Zero"
	}

	m := 1000000000

	for m >= 1 {
		a := num / m
		b := num % m
		if a != 0 {
			res += convertTrunk(a)
			switch m {
			case 1000000000:
				res += " Billion "
			case 1000000:
				res += " Million "
			case 1000:
				res += " Thousand "
			}
		}
		if b != 0 {
			num = b
		}

		m /= 1000
	}

	return res
}

func convertTrunk(num int) string {
	if num < 0 || num > 999 {
		return ""
	}

	var res string
	if _, ok := dict[num]; ok {
		return dict[num]
	} else {

	}

	a := num / 100
	b := num % 100
	if a > 0 {
		res += dict[a] + " Hundred"
	}
	if b > 0 {
		if _, ok := dict[b]; ok {
			res += " " + dict[b]
			return res
		} else {
			a = b / 10
			b = b % 10
			res += " " + dict[a*10] + " " + dict[b]
		}
	}

	return res
}
