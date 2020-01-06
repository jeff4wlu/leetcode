package leet73

func PalindromePairs(in []string) [][]int {

	n := len(in)

	res := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if in[i] == in[j] {
				continue
			}
			if isPalindrome(in[i] + in[j]) {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func isPalindrome(in string) bool {

	n := len(in)

	for i := 0; i < (n+1)/2; i++ {
		if in[i:i+1] != in[n-1-i:n-i] {
			return false
		}
	}

	return true
}
