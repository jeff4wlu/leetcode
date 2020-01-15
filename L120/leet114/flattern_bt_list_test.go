package leet114

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlatternBT2List(t *testing.T) {

	Convey("TestFlatternBT2List", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 5

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 3
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 4

			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 6
			FlatternBT2List(tmp)
			got := infra.GetInorderBT(tmp)
			want := []int{1, 2, 3, 4, 5, 6}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
