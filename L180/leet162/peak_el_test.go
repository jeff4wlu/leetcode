package leet162

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPeakEl(t *testing.T) {

	Convey("TestPeakEl", t, func() {
		Convey("1", func() {

			got := PeakEl([]int{1, 2, 3, 1})
			want := 2

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := PeakEl([]int{1, 2, 1, 3, 5, 6, 4})

			if !(got == 1 || got == 5) {
				t.Errorf("failed")
			}

		})

	})

}
