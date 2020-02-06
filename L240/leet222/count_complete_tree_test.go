package leet222

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountCompleteTree(t *testing.T) {

	Convey("TestCountCompleteTree", t, func() {
		Convey("相同", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			tmp.Left.Left = &infra.BTIntNode{}
			tmp.Left.Left.Value = 4
			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5

			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 6

			got := CountCompleteTree(tmp)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
