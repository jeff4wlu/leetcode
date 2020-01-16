package leet53

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSubarr(t *testing.T) {

	Convey("TestMaxSubarr", t, func() {
		Convey("1", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			got := MaxSubarr1(nums)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

	Convey("TestMaxSubarr2", t, func() {
		Convey("1", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			got := MaxSubarr2(nums)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

	Convey("TestMaxSubarr3", t, func() {
		Convey("1", func() {
			nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
			got := MaxSubarr3(nums)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
