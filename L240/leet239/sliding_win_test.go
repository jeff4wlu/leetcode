package leet239

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSlidingWin(t *testing.T) {

	Convey("TestSlidingWin", t, func() {
		Convey("用例1", func() {

			got := SlidingWin([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
			want := []int{3, 3, 5, 5, 6, 7}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
