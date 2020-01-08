package leet81

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchRotateSortedArr(t *testing.T) {

	Convey("TestSearchRotateSortedArr", t, func() {
		Convey("用例1", func() {

			got := SearchRotateSortedArr([]int{2, 5, 6, 0, 0, 1, 2}, 0)
			want := true

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("用例2", func() {

			got := SearchRotateSortedArr([]int{2, 5, 6, 0, 0, 1, 2}, 3)
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

		Convey("用例3", func() {

			got := SearchRotateSortedArr([]int{2, 5, 6, 0, 0, 1, 2}, 7)
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}

		})

	})

}
