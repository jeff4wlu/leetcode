package leet201

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumRange(t *testing.T) {

	Convey("TestNumRange", t, func() {
		Convey("用例1", func() {

			got := NumRange(5, 7)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := NumRange(0, 1)
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例3", func() {

			got := NumRange(26, 30)
			want := 24

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
