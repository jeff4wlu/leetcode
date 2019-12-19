package leet22

//left,还剩下多少个左括号
//right,还剩下多少个右括号
//word，当前的word是什么
//res，结果集
func GenParenthesesDFS(left, right int, word string, res *[]string) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, word)
	}

	if left > 0 {
		GenParenthesesDFS(left-1, right, word+"(", res)
	}
	if right > 0 {
		GenParenthesesDFS(left, right-1, word+")", res)
	}

}

// append作用于slice前是不需要用make分配内存
func GenParentheses(n int) []string {
	//res := make([]string, 10)
	var res []string
	GenParenthesesDFS(n, n, "", &res)
	return res
}
