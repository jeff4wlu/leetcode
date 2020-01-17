package leet124

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTMaxPath(t *testing.T) {

	Convey("TestBTMaxPath", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = -1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = -2
			//tmp.Right = &infra.BTIntNode{}
			//tmp.Right.Value = 3
			/*
				tmp.Right.Left = &infra.BTIntNode{}
				tmp.Right.Left.Value = 15
				tmp.Right.Right = &infra.BTIntNode{}
				tmp.Right.Right.Value = 7*/

			got := BTMaxPath(tmp)
			want := -1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3
			/*
				tmp.Right.Left = &infra.BTIntNode{}
				tmp.Right.Left.Value = 15
				tmp.Right.Right = &infra.BTIntNode{}
				tmp.Right.Right.Value = 7*/

			got := BTMaxPath(tmp)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例3", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = -10
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 9
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 20

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 15
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 7

			got := BTMaxPath(tmp)
			want := 42

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
