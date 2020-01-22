package leet150

import (
	"github.com/golang-collections/collections/stack"
	"strconv"
)

func EvalRPN(in []string) int {

	stack := stack.New()

	for _, v := range in {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			a, _ := strconv.Atoi(stack.Pop().(string))
			b, _ := strconv.Atoi(stack.Pop().(string))

			var c int
			switch v {
			case "+":
				c = b + a
			case "-":
				c = b - a
			case "*":
				c = b * a
			case "/":
				c = b / a
			}
			stack.Push(strconv.Itoa(c))

		} else {
			stack.Push(v)
		}

	}
	res, _ := strconv.Atoi(stack.Pop().(string))

	return res
}
