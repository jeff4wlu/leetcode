package leet135

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCandy(t *testing.T) {

	Convey("TestCandy", t, func() {
		Convey("用例1", func() {

			got := Candy([]int{1, 0, 2})
			want := 5

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := Candy([]int{1, 2, 2})
			want := 4

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
