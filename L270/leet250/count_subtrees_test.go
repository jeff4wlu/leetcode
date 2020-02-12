package leet250

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountSubtrees(t *testing.T) {

	Convey("TestCountSubtrees", t, func() {
		Convey("相同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 5
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 5

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 5
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5

			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 5

			got := CountSubtrees(tmp)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
