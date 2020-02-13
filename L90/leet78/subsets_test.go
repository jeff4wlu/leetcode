package leet78

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsets(t *testing.T) {

	Convey("TestSubsets", t, func() {
		Convey("用例1", func() {

			got := Subsets([]int{1, 2, 3})
			want := [][]int{
				{3},
				{1},
				{2},
				{1, 2, 3},
				{1, 3},
				{2, 3},
				{1, 2},
				{},
			}

			if got == nil || !infra.IntarrCollectionComp(got, want, true) {
				t.Errorf("failed")
			}

		})

	})

}
