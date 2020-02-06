package leet214

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortestPalindrome(t *testing.T) {

	Convey("TestShortestPalindrome", t, func() {
		Convey("常规", func() {

			got := ShortestPalindrome("aacecaaa")
			want := "aaacecaaa"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("常规1", func() {

			got := ShortestPalindrome("abcd")
			want := "dcbabcd"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("常规2", func() {

			got := ShortestPalindrome("aa")
			want := "aa"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("常规3", func() {

			got := ShortestPalindrome("ab")
			want := "bab"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}
