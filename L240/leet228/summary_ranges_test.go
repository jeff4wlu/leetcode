package leet228

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSummaryRanges(t *testing.T) {

	Convey("TestSummaryRanges", t, func() {
		Convey("用例1", func() {

			got := SummaryRanges([]int{0, 1, 2, 4, 5, 7})
			want := []string{"0->2", "4->5", "7"}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
			want := []string{"0", "2->4", "6", "8->9"}

			if !infra.StringArrSeqComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
