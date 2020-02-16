package leet270

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
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 3

			got := ClosestValue(tmp, 3.714286)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
