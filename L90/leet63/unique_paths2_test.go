package leet63

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniquePaths2(t *testing.T) {

	Convey("TestUniquePaths2", t, func() {
		Convey("用例1", func() {
			grid := [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}

			_, got := UniquePaths2(grid)
			want := [][]string{{"right", "right", "down", "down"}, {"down", "down", "right", "right"}}

			if !infra.StringarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
