package leet112

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPathSum(t *testing.T) {

	Convey("TestPathSum", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 4
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 8

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 11

			tmp.Left.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Left.Value = 7
			tmp.Left.Left.Right = &infra.BTIntNode{}
			tmp.Left.Left.Right.Value = 2

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 13
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 4

			tmp.Right.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Right.Value = 1

			got := PathSum(tmp, 22)
			want := true

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

	})

}
