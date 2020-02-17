package leet278

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindBadVer(t *testing.T) {

	Convey("TestFindBadVer", t, func() {

		Convey("1", func() {

			got := FindBadVer([]int{0, 0, 0, 0, 1, 1})
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := FindBadVer([]int{0, 1, 1, 1, 1, 1})
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
