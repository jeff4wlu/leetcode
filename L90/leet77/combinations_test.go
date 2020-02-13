package leet77

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinations(t *testing.T) {

	Convey("TestCombinations", t, func() {
		Convey("用例1", func() {

			got := Combinations(4, 2)
			want := [][]int{
				{2, 4},
				{3, 4},
				{2, 3},
				{1, 2},
				{1, 3},
				{1, 4},
			}

			if got == nil || !infra.IntarrCollectionComp(got, want, true) {
				t.Errorf("failed")
			}

		})

	})

}
