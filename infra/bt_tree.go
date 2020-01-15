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

//获取后序数组
func GetPosBT(tn *BTIntNode) []int {
	path := []int{}
	getPostBT(tn, &path)
	return path

}

func getPostBT(tn *BTIntNode, path *[]int) {

	if tn == nil {
		return
	}

	getPostBT(tn.Left, path)
	getPostBT(tn.Right, path)
	*path = append(*path, tn.Value)
}

//获取中序数组
func GetInorderBT(tn *BTIntNode) []int {
	path := []int{}
	getInorderBT(tn, &path)
	return path
}

func getInorderBT(tn *BTIntNode, path *[]int) {

	if tn == nil {
		return
	}

	getInorderBT(tn.Left, path)
	*path = append(*path, tn.Value)
	getInorderBT(tn.Right, path)

}
