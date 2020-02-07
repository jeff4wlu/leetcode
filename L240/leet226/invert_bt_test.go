package leet226

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvertBT(t *testing.T) {

	Convey("TestInvertBT", t, func() {
		Convey("相同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 4
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 7

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 1
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 6
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 9

			got := InvertBT(tmp)

			infra.PrintTree(got, t)

		})

	})

	Convey("TestInvertBTDFS", t, func() {
		Convey("相同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 4
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 7

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 1
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 6
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 9

			got := InvertBTDFS(tmp)

			infra.PrintTree(got, t)

		})

	})

}
