package leet110

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBalancedBT(t *testing.T) {

	Convey("TestBalancedBT", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 3
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 9
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 20

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 15
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 7

			got := BalancedBT(tmp)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 2

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 3
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			tmp.Left.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Left.Value = 4
			tmp.Left.Left.Right = &infra.BTIntNode{}
			tmp.Left.Left.Right.Value = 4

			got := BalancedBT(tmp)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
