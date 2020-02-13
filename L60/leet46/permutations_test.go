package leet46

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermutations(t *testing.T) {

	Convey("TestPermutations", t, func() {
		Convey("用例1", func() {
			nums := []int{1, 2, 3}

			res := Permutations(nums)
			got := [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

	})

}
