package leet246

func StrobogrammaticNum(word string) bool {

	n := len(word)

	for i := 0; i < n; i++ {

		if word[i] == word[n-i-1] && (word[i] == '0' || word[i] == '1' || word[i] == '8') {
			continue
		} else if word[i] != word[n-i-1] && (word[i] == '6' || word[i] == '9' || word[n-i-1] == '6' || word[n-i-1] == '9') {
			continue
		}
		return false
	}
	return true
}
