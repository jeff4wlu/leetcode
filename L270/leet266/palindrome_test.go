package leet266

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindrome(t *testing.T) {

	Convey("TestPalindrome", t, func() {

		Convey("1", func() {

			got := Palindrome("code")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := Palindrome("aab")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := Palindrome("carerac")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
