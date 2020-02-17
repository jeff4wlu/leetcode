package leet275

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHIndex(t *testing.T) {

	Convey("TestHIndex", t, func() {

		Convey("1", func() {

			got := HIndex([]int{0, 1, 3, 5, 6})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := HIndex([]int{0, 1, 2, 5, 6})
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("3", func() {

			got := HIndex([]int{0, 1, 4, 5, 6})
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
