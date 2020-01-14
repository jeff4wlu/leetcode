package infra

import (
	"fmt"
	"testing"
)

type BTIntNode struct {
	Value int
	Left  *BTIntNode
	Right *BTIntNode
}

//这个是前序打印
func PrintTree(tnode *BTIntNode, t *testing.T) {

	if tnode == nil {
		fmt.Printf(" null")
		return
	}
	fmt.Printf(" %d", tnode.Value)
	PrintTree(tnode.Left, t)
	PrintTree(tnode.Right, t)
}
