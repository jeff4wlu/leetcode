package leet9

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindromeNum(t *testing.T) {

	Convey("TestPalindromeNum", t, func() {
		Convey("负数", func() {
			in := -121
			got := PalindromeNum(in)
			want := false

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey("回文", func() {
			in := 121
			got := PalindromeNum(in)
			want := true

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

		Convey("非回文", func() {
			in := 10
			got := PalindromeNum(in)
			want := false

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}

		})

	})

}
