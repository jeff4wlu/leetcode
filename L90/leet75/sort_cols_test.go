package leet75

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortCols(t *testing.T) {

	Convey("TestSortCols", t, func() {
		Convey("用例1", func() {

			got := []int{2, 0, 2, 1, 1, 0}
			SortCols(&got)
			want := []int{0, 0, 1, 1, 2, 2}

			if !infra.IntArrCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
