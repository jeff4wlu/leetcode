package leet156

import (
	"fmt"
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTUpsideDown(t *testing.T) {

	Convey("TestBTUpsideDown", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3
			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 4
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5
			fmt.Println()
			infra.PrintTree(tmp, t)
			fmt.Println()

			res := BTUpsideDown(tmp)
			fmt.Println()
			infra.PrintTree(res, t)
			fmt.Println()

		})

	})

}
