package leet215

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKthLargest(t *testing.T) {

	Convey("TestKthLargest", t, func() {
		Convey("无重复数", func() {

			got := KthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("有重复数", func() {

			got := KthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("有重复数，并刚好在重复数上1", func() {

			got := KthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("有重复数，并刚好在重复数上2", func() {

			got := KthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 2)
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
