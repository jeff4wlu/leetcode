package leet32

func LongestValidParenthese(s string) int {

	n := len(s)

	var left, right, longest int

	for j := 0; j < n; j++ {
		for i := j; i < n; i++ {
			if s[i:i+1] == "(" {
				left++
			} else {
				right++
			}

			if right > left {
				if i-j > longest {
					longest = i - j
				}
				break
			}

		}
		left = 0
		right = 0
	}

	for j := n - 1; j >= 0; j-- {
		for i := j; i >= 0; i-- {
			if s[i:i+1] == "(" {
				left++
			} else {
				right++
			}

			if right < left {
				if j-i > longest {
					longest = j - i
				}
				break
			}

		}
		left = 0
		right = 0
	}

	return longest

}
