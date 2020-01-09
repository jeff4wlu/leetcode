package leet90

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsets(t *testing.T) {

	Convey("TestSubsets", t, func() {
		Convey("用例1", func() {

			got := Subsets([]int{1, 2, 2})
			want := [][]int{
				{2},
				{1},
				{1, 2, 2},
				{2, 2},
				{1, 2},
				{},
			}

			if got == nil || !infra.IntarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
