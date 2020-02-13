package leet257

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTPath(t *testing.T) {

	Convey("TestBTPath", t, func() {
		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 2
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 3

			tmp.Left.Right = &infra.BTIntNode{}
			tmp.Left.Right.Value = 5

			got := BTPath(tmp)
			want := []string{
				"1->2->5", "1->3",
			}
			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
