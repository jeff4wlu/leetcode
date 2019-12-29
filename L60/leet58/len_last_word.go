package leet58

func LenLastWord(s string) int {

	n := len(s)
	if n <= 0 {
		return 0
	}

	var tmp string
	for i := n - 1; i >= 0; i-- {
		if s[i:i+1] != " " {
			tmp = s[:i+1]
			break
		}
	}

	n = len(tmp)

	for i := n - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			return n - i - 1
		}
	}

	return 0
}
