package leet281

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestZigzagIterator(t *testing.T) {

	Convey("TestZigzagIterator", t, func() {

		Convey("1", func() {

			nums := [][]int{
				{1, 2, 3},
				{4, 5, 6, 7},
				{8, 9},
			}

			got := ZigzagIterator(nums)
			want := []int{1, 4, 8, 2, 5, 9, 3, 6, 7}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
