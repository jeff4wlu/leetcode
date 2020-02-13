package leet113

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPathSum2(t *testing.T) {

	Convey("TestPathSum2", t, func() {
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

			tmp.Right.Right.Left = &infra.BTIntNode{}
			tmp.Right.Right.Left.Value = 5
			tmp.Right.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Right.Value = 1

			got := PathSum2(tmp, 22)
			want := [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			}

			if got == nil || !infra.IntarrCollectionComp(got, want, true) {
				t.Errorf("failed")
			}

		})

	})

}
