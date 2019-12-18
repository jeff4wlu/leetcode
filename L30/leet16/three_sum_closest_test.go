package leet16

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestThreeSumClosest(t *testing.T) {

	Convey("TestThreeSumClosest", t, func() {
		Convey("常规", func() {
			arr := []int{-1, 2, 1, -4}
			got := ThreeSumClosest(arr, 1)
			want := 2

			if got != want {
				t.Errorf("got  %d, want  %d", got, want)
			}

		})

		Convey("只有3个数的时候", func() {
			arr := []int{-1, 2, 1}
			got := ThreeSumClosest(arr, 1)
			want := 2

			if got != want {
				t.Errorf("got  %d, want  %d", got, want)
			}

		})

	})

}
