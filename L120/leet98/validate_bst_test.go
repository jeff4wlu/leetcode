package leet98

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateBST(t *testing.T) {

	Convey("TestValidateBST", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 2
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			got := ValidateBST(tmp)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 4
			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 3
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 6

			got := ValidateBST(tmp)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestValidateBST1", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 2
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			got := ValidateBST1(tmp)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 4
			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 3
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 6

			got := ValidateBST1(tmp)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
