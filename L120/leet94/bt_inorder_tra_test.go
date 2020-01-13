package leet94

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTInorderTra(t *testing.T) {

	Convey("TestBTInorderTra", t, func() {
		Convey("用例1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 1
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 2
			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 3

			got := BTInorderTra(tmp)
			want := []int{1, 3, 2}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
