package leet87

//递归，还可以使用DP
func ScrambleStr(a, b string) bool {

	return solve(a, b)
}

func solve(a, b string) bool {

	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return false
	}

	if a == b {
		return true
	}

	for i := 1; i < aLen; i++ {
		if solve(a[0:i], b[0:i]) && solve(a[i:], b[i:]) || solve(a[0:i], b[i:]) && solve(a[i:], b[0:i]) {
			return true
		}
	}
	return false
}
