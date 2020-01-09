package leet88

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeSortedArr(t *testing.T) {

	Convey("TestMergeSortedArr", t, func() {
		Convey("用例1", func() {

			got := []int{1, 2, 3, 0, 0, 0}
			MergeSortedArr(got, []int{2, 5, 6}, 3, 3)
			want := []int{1, 2, 2, 3, 5, 6}

			if !infra.IntArrCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := []int{1, 2, 6, 0, 0, 0}
			MergeSortedArr(got, []int{2, 5, 6}, 3, 3)
			want := []int{1, 2, 2, 5, 6, 6}

			if !infra.IntArrCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
