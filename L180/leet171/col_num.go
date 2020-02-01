package leet171

func ColNum(s string) int {

	n := len(s)
	if n <= 0 || n > 2 {
		return -1
	}

	if n == 1 {
		return int(s[0]) - int('A') + 1
	}

	return (int(s[0])-int('A')+1)*26 + (int(s[1]) - int('A') + 1)
}
