package leet39

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinationSum(t *testing.T) {

	Convey("TestCombinationSum", t, func() {
		Convey("用例1", func() {
			nums := []int{2, 3, 6, 7}

			res := CombinationSum(nums, 7)
			got := [][]int{{2, 2, 3}, {7}}

			if res == nil || !infra.IntarrCollectionComp(res, got, false) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{2, 3, 5}

			res := CombinationSum(nums, 8)
			got := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

			if res == nil || !infra.IntarrCollectionComp(res, got, false) {
				t.Errorf("failed")
			}

		})

		Convey("边界", func() {
			nums := []int{2, 3, 5}

			res := CombinationSum(nums, 1)
			//var got [][]int

			if len(res) != 0 {
				t.Errorf("failed")
			}

		})

	})

}
