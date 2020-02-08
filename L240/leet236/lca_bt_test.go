package leet236

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLCABt(t *testing.T) {

	Convey("TestLCABt", t, func() {
		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 3
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 5
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 1

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 6
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 2

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 0
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 8

			tmp.Left.Right.Left = &infra.BTIntNode{}
			tmp.Left.Right.Left.Value = 7
			tmp.Left.Right.Right = &infra.BTIntNode{}
			tmp.Left.Right.Right.Value = 4

			got := LCABt(tmp, 5, 1)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 3
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 5
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 1

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 6
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 2

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 0
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 8

			tmp.Left.Right.Left = &infra.BTIntNode{}
			tmp.Left.Right.Left.Value = 7
			tmp.Left.Right.Right = &infra.BTIntNode{}
			tmp.Left.Right.Right.Value = 4

			got := LCABt(tmp, 5, 4)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
