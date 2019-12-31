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

	Convey("TestUniquePaths22", t, func() {
		Convey("用例1", func() {
			grid := [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			}

			got := UniquePaths22(grid)
			want := 1

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			grid := [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}

			got := UniquePaths22(grid)
			want := 2

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {
			grid := [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{1, 0, 0},
			}

			got := UniquePaths22(grid)
			want := 0

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
