package leet33

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchRotatedSortedArr(t *testing.T) {

	Convey("TestSearchRotatedSortedArr", t, func() {
		Convey("用例1", func() {
			arr := []int{4, 5, 6, 7, 0, 1, 2}
			got := SearchRotatedSortedArr(arr, 0)
			want := 4

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("用例2", func() {
			arr := []int{4, 5, 6, 7, 0, 1, 2}
			got := SearchRotatedSortedArr(arr, 3)
			want := -1

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
