package leet216

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombinationSum3(t *testing.T) {

	Convey("TestCombinationSum3", t, func() {
		Convey("1", func() {

			got := CombinationSum3(3, 7)
			want := [][]int{
				{1, 2, 4},
			}

			if !infra.IntarrCollectionComp(got, want, true) {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := CombinationSum3(3, 9)
			want := [][]int{
				{1, 2, 6}, {1, 3, 5}, {2, 3, 4},
			}

			if !infra.IntarrCollectionComp(got, want, true) {
				t.Errorf("failed")
			}

		})

	})

}
