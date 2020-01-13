package leet95

import (
	"fmt"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueBST(t *testing.T) {

	Convey("TestUniqueBST", t, func() {
		Convey("用例1", func() {

			got := UniqueBST(3)
			fmt.Println()
			for _, v := range got {
				printTree(v, t)
				fmt.Println()
			}

		})

	})

}

func printTree(tnode *infra.BTIntNode, t *testing.T) {

	if tnode == nil {
		fmt.Printf(" null")
		return
	}
	fmt.Printf(" %d", tnode.Value)
	printTree(tnode.Left, t)
	printTree(tnode.Right, t)
}
