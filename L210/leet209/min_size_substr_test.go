package leet209

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinSizeSubstr(t *testing.T) {

	Convey("TestMinSizeSubstr", t, func() {
		Convey("常规", func() {
			got := MinSizeSubstr([]int{2, 3, 1, 2, 4, 3}, 7)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("常规1", func() {
			got := MinSizeSubstr([]int{2}, 7)
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("常规2", func() {
			got := MinSizeSubstr([]int{2}, 1)
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
