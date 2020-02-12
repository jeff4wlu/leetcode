package leet251

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlattern2D(t *testing.T) {

	Convey("TestFlattern2D", t, func() {
		Convey("相同", func() {

			nums := [][]int{
				{1, 2},
				{3},
				{4, 5, 6},
			}

			got := Flattern2D(nums)
			want := []int{1, 2, 3, 4, 5, 6}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("1", func() {

			nums := [][]int{
				{},
				{1, 2},
				{3},
				{},
				{4, 5, 6},
				{},
			}

			got := Flattern2D(nums)
			want := []int{1, 2, 3, 4, 5, 6}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
