package leet103

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTZLvlOrderTra(t *testing.T) {

	Convey("TestBTZLvlOrderTra", t, func() {
		Convey("用例1", func() {
			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 3
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 9
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 20

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 15
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 7

			got := BTZLvlOrderTra(tmp)
			want := [][]int{
				{3},
				{20, 9},
				{15, 7},
			}

			if !infra.IntArr2DComp(want, got) {
				t.Errorf("failed")
			}

		})

	})

}
