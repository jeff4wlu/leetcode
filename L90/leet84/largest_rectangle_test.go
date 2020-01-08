package leet84

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLargestRect(t *testing.T) {

	Convey("TestLargestRect", t, func() {
		Convey("用例1", func() {

			got := LargestRect([]int{2, 1, 5, 6, 2, 3})
			want := 10

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := LargestRect([]int{20, 1, 5, 6, 2, 3})
			want := 20

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

	Convey("TestLargestRect1", t, func() {
		Convey("用例1", func() {

			got := LargestRect1([]int{2, 1, 5, 6, 2, 3})
			want := 10

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := LargestRect1([]int{20, 1, 5, 6, 2, 3})
			want := 20

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
