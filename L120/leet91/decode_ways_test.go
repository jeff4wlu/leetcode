package leet91

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeWays(t *testing.T) {

	Convey("TestDecodeWays", t, func() {
		Convey("用例1", func() {

			got := DecodeWays("12")
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := DecodeWays("226")
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
