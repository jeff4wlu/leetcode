package leet199

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTRightView(t *testing.T) {

	Convey("TestBTRightView", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 4

			got := BTRightView(tmp)
			want := []int{1, 3, 4}

			if !infra.IntArrSeqCmp(got, want) {
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
			tmp.Right.Value = 3

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5

			got := BTRightView(tmp)
			want := []int{1, 3, 5}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
