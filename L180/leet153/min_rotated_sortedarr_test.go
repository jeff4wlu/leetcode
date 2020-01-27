package leet153

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPro(t *testing.T) {

	Convey("TestMinRotatedSortedArr", t, func() {
		Convey("1", func() {
			nums := []int{3, 4, 5, 1, 2}
			got := MinRotatedSortedArr(nums)
			want := 1

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("2", func() {
			nums := []int{4, 5, 6, 7, 0, 1, 2}
			got := MinRotatedSortedArr(nums)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("重复值", func() {
			nums := []int{0, 1, 2, 2, 2, 0}
			got := MinRotatedSortedArr(nums)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
