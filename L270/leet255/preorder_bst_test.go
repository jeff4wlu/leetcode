package leet255

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreorderBST(t *testing.T) {

	Convey("TestPreorderBST", t, func() {
		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 6

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 1
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			got := PreorderBST(tmp, []int{5, 2, 6, 1, 3})
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 6

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 1
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			got := PreorderBST(tmp, []int{5, 2, 1, 3, 6})
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
