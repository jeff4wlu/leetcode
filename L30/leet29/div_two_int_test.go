package leet29

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDivTwoInt(t *testing.T) {

	Convey("TestDivTwoInt", t, func() {
		Convey("常规1", func() {

			got := DivTwoInt(10, 3)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("负数", func() {

			got := DivTwoInt(7, -3)
			want := -2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("被除数为0", func() {

			got := DivTwoInt(0, -3)
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
