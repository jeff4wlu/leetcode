package leet168

func ColTitle(col int) string {
	if col < 1 {
		return ""
	}
	if col <= 26 {
		res := 'A' + col - 1
		return string(res)
	}
	a := col / 26
	b := col % 26

	if a > 26 || b > 26 {
		return ""
	}

	return string('A'+a-1) + string('A'+b-1)

}
