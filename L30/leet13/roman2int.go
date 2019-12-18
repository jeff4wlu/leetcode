package leet13

var dict = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {

	var res int

	for i := 0; i < len(s); i++ {
		val := dict[s[i:i+1]]
		if i == len(s)-1 || dict[s[i+1:i+2]] <= dict[s[i:i+1]] {
			res += val
		} else {
			res -= val
		}

	}

	return res

}
