package leet69

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSqrt(t *testing.T) {

	Convey("TestSqrt", t, func() {
		Convey("用例1", func() {

			got := Sqrt(4)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := Sqrt(8)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
