package leet111

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinBTDepth(t *testing.T) {

	Convey("TestMinBTDepth", t, func() {
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

			got := MinBTDepth(tmp)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
