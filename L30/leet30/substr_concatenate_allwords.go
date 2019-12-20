package leet30

func SubstrConcatAllWords(s string, w []string) []int {

	words := len(w)
	word := len(w[0])
	total := words * word
	stringLen := len(s)

	var res []int

	dict := make(map[string]int, words)

	for i := 0; i < words; i++ {
		dict[w[i]] = 0
	}

	for i := 0; i < stringLen; i++ {
		tmp := i
		isOk := true
		for j := 0; j < words; j++ {
			if _, ok := dict[s[tmp:tmp+word]]; ok {
				dict[s[tmp:tmp+word]]++
				if dict[s[tmp:tmp+word]] > 1 {
					isOk = false
					break
				}
			} else {
				isOk = false
				break
			}
			tmp += word
		}

		if isOk {
			res = append(res, i)
		}

		if stringLen-i < total {
			break
		}

		for k := range dict {
			dict[k] = 0
		}

	}

	return res

}
