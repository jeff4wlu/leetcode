package leet129

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSumRootLeafNum(t *testing.T) {

	Convey("TestSumRootLeafNum", t, func() {
		Convey("用例1", func() {

			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			got := SumRootLeafNum(tmp)
			want := 25

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			tmp := &infra.BTIntNode{}
			tmp.Value = 4
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 9
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 0

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 5
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 1

			got := SumRootLeafNum(tmp)
			want := 1026

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
