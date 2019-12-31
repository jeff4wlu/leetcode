package leet64

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinPathSum(t *testing.T) {

	Convey("TestMinPathSum", t, func() {
		Convey("用例1", func() {
			grid := [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}

			got := MinPathSum(grid)
			want := 7

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
