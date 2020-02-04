package leet198

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHouseRobber(t *testing.T) {

	Convey("TestHouseRobber", t, func() {
		Convey("用例1", func() {

			got := HouseRobber([]int{1, 2, 3, 1})
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := HouseRobber([]int{2, 7, 9, 3, 1})
			want := 12

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
