package leet249

var dict map[string][]string //全局变量声明与初始化

func init() {
	dict = map[string][]string{}
}

//找规律，字母循环规律，和利用map
func GroupShifted(words []string) [][]string {

	for _, v := range words {
		word := getBaseWord(v)
		if _, ok := dict[word]; !ok {
			dict[word] = []string{}
		}
		dict[word] = append(dict[word], v)
	}

	res := [][]string{}
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

func getBaseWord(word string) string {
	l := len(word)
	gap := int(word[0] - 'a')
	res := []byte{}
	for i := 0; i < l; i++ {
		tmp := word[i] - byte(gap)
		if tmp < 'a' {
			tmp += 26
		}
		res = append(res, tmp)
	}

	return string(res)

}
