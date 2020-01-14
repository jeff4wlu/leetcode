package leet101

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSymetricTree(t *testing.T) {

	Convey("TestSymetricTree", t, func() {
		Convey("是", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 2

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 3
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 3

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 4
			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 4

			got := SymetricTree(tmp)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("不是", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 2

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 3

			got := SymetricTree(tmp)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
