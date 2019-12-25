package leet47

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutations2(t *testing.T) {

	Convey("TestPermutations2", t, func() {
		Convey("用例1", func() {
			nums := []int{1, 1, 2}

			res := Permutations2(nums)
			got := [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := []int{1, 1, 2, 2}

			res := Permutations2(nums)
			got := [][]int{
				{1, 1, 2, 2},
				{1, 2, 1, 2},
				{2, 1, 1, 2},
				{1, 2, 2, 1},
				{2, 2, 1, 1},
				{2, 1, 2, 1},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got) {
				t.Errorf("failed")
			}

		})

	})

}
