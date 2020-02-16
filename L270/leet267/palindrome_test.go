package leet267

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindrome(t *testing.T) {

	Convey("TestPalindrome", t, func() {

		Convey("回文偶数", func() {

			got := Palindrome("aabb")
			want := []string{"abba", "baab"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("非回文", func() {

			got := Palindrome("code")
			want := []string{}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("一个字符", func() {

			got := Palindrome("c")
			want := []string{"c"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("回文基数", func() {

			got := Palindrome("aaabb")
			want := []string{"ababa", "baaab"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})
		Convey("有重复字符的回文", func() {

			got := Palindrome("aabaaab")
			want := []string{"aababaa", "abaaaba", "baaaaab"}

			if !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})
	})
}
