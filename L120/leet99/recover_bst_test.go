package leet99

import (
	"fmt"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRecoverBST(t *testing.T) {

	Convey("TestRecoverBST", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 3
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 2
			fmt.Println()
			infra.PrintTree(tmp, t)
			fmt.Println()

			RecoverBST(tmp)
			fmt.Println()
			infra.PrintTree(tmp, t)
			fmt.Println()

		})

	})

}
