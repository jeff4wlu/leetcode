package leet56

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeInterval(t *testing.T) {

	Convey("TestMergeInterval", t, func() {
		Convey("用例1", func() {
			nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

			res := MergeInterval(nums)
			got := [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

		Convey("重叠", func() {
			nums := [][]int{{1, 4}, {3, 5}}

			res := MergeInterval(nums)
			got := [][]int{
				{1, 5},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

		Convey("全包含", func() {
			nums := [][]int{{1, 4}, {2, 3}, {9, 11}}

			res := MergeInterval(nums)
			got := [][]int{
				{1, 4},
				{9, 11},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

	})

}
