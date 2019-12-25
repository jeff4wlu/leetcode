package leet40

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinationSumUniq(t *testing.T) {

	Convey("TestCombinationSumUniq", t, func() {
		Convey("用例1", func() {
			nums := []int{10, 1, 2, 7, 6, 1, 5}

			res := CombinationSumUniq(nums, 8)
			got := [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}

			if res == nil || !infra.IntarrCollectionComp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{2, 5, 2, 1, 2}

			res := CombinationSumUniq(nums, 5)
			got := [][]int{{5}, {1, 2, 2}}

			if res == nil || !infra.IntarrCollectionComp(res, got) {
				t.Errorf("failed")
			}

		})

	})

}
