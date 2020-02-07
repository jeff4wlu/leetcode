package leet230

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKthElBT(t *testing.T) {

	Convey("TestKthElBT", t, func() {
		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 3
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 4

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 2

			got := KthElBT(tmp, 1)
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 3
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 6

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 2
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 4

			tmp.Left.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Left.Value = 1

			got := KthElBT(tmp, 3)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
