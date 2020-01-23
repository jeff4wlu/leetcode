package leet151

import "github.com/golang-collections/collections/stack"

func ReverseWord(s string) string {

	n := len(s)
	if n < 1 {
		return ""
	}

	stack := stack.New()

	i := 0
	var start int
	for i < n {

		for ; i < n; i++ {
			if s[i] != ' ' {
				break
			}
		}
		if i == n {
			if stack.Len() == 0 {
				return ""
			}
			break
		}
		start = i
		for ; i < n; i++ {
			if s[i] == ' ' {
				stack.Push(s[start:i])
				break
			}
		}
		if i == n {
			stack.Push(s[start:])
			break
		}

	}

	if stack.Len() == 0 {
		return ""
	}

	var res string
	for stack.Len() != 0 {
		res += stack.Pop().(string)
		res += " "
	}

	res = res[:len(res)-1]
	return res
}
