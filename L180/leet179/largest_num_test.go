package leet179

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLargestNum(t *testing.T) {

	Convey("TestLargestNum", t, func() {
		Convey("用例1", func() {

			got := LargestNum([]int{10, 2})
			want := "210"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("用例2", func() {

			got := LargestNum([]int{3, 30, 34, 5, 9})
			want := "9534330"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}
