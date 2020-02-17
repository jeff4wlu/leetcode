package leet272

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClosestValue(t *testing.T) {

	Convey("TestClosestValue", t, func() {

		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 4
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 5

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 1
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 3

			got := ClosestVal(tmp, 3.714286, 3)
			want := []int{4, 3, 5}

			if !infra.IntArrNonSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
