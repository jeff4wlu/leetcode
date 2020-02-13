package leet57

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertInterval(t *testing.T) {

	Convey("TestInsertInterval", t, func() {
		Convey("用例1", func() {
			nums := [][]int{{1, 3}, {6, 9}}

			res := InsertInterval(nums, []int{2, 5})
			got := [][]int{
				{1, 5},
				{6, 9},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			nums := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}

			res := InsertInterval(nums, []int{4, 8})
			got := [][]int{
				{1, 2}, {3, 10}, {12, 16},
			}

			if res == nil || !infra.IntarrCollectionComp(res, got,true) {
				t.Errorf("failed")
			}

		})

	})

}
