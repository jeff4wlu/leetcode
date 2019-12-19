package leet20

import "leetcode/infra"

var stack infra.Stack = infra.CreateStack()

var dict = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func ValidParentheses(s string) bool {
	size := len(s)
	if size < 2 || size > 2 && size%2 != 0 {
		return false
	}

	for i := 0; i < size; i++ {
		char := s[i : i+1]

		if _, ok := dict[char]; !ok { //入栈左边
			stack.Push(char)
		} else {
			tmp := stack.Pop()
			if tmp == nil || tmp != dict[char] {
				return false
			}
		}
	}

	return true

}
