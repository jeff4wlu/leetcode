package leet70

import "leetcode/infra"

func SimplifyPath(path string) string {

	n := len(path)

	stack := infra.CreateStack()

	var tokenStart, tokenEnd int
	for i := 0; i < n; i++ {
		tokenEnd = i
		tmp := path[i : i+1]
		if tmp == "/" && tokenEnd-tokenStart == 1 { //连续斜杠
			tokenStart = tokenEnd
			continue
		}
		if tmp == "/" && tokenEnd-tokenStart > 1 { //有效token

			token := path[tokenStart : tokenEnd+1]
			if token == "/../" {
				stack.Pop()
			} else if token != "/./" {
				stack.Push(path[tokenStart:tokenEnd])
			}

			tokenStart = tokenEnd
		}
	}

	if stack.IsEmpty() {
		return "/"
	}

	stack2 := infra.CreateStack()
	for {
		stack2.Push(stack.Pop())
		if stack.IsEmpty() {
			break
		}
	}

	var res string
	for {
		res += (stack2.Pop()).(string)
		if stack2.IsEmpty() {
			break
		}
	}

	return res
}
