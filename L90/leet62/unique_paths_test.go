package leet62

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniquePaths(t *testing.T) {

	Convey("TestUniquePaths", t, func() {
		Convey("用例1", func() {

			_, got := UniquePaths(2, 3)
			want := [][]string{{"right", "right", "down"}, {"right", "down", "right"}, {"down", "right", "right"}}

			if !infra.StringArrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got, _ := UniquePaths(7, 3)
			want := 28

			if got != want {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestUniquePaths2", t, func() {
		Convey("用例1", func() {

			got := UniquePaths2(2, 3)
			want := 3

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := UniquePaths2(7, 3)
			want := 28

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
