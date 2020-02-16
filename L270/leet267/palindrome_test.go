package leet267

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindrome(t *testing.T) {

	Convey("TestPalindrome", t, func() {

		Convey("1", func() {

			got := Palindrome("aabb")
			want := []string{"abba", "baab"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})
	})
}
