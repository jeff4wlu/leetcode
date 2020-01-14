package leet100

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSameTree(t *testing.T) {

	Convey("TestSameTree", t, func() {
		Convey("相同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 3
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 2

			tmp1 := &infra.BTIntNode{}
			tmp1.Value = 1
			tmp1.Left = &infra.BTIntNode{}
			tmp1.Left.Value = 3
			tmp1.Left.Right = &infra.BTIntNode{}
			tmp1.Left.Right.Value = 2

			got := SameTree(tmp, tmp1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("不同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 3
			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 2

			tmp1 := &infra.BTIntNode{}
			tmp1.Value = 1
			tmp1.Left = &infra.BTIntNode{}
			tmp1.Left.Value = 3
			tmp1.Left.Right = &infra.BTIntNode{}
			tmp1.Left.Right.Value = 2

			got := SameTree(tmp, tmp1)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
