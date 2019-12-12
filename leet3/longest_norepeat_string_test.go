package leet3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	Convey("TestLengthOfLongestSubstring", t, func() {
		Convey("求最长的无重复字符串长度", func() {
			var a = []byte("helloboy")

			got := LengthOfLongestSubstring(a)
			want := 3

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("重复字符串", func() {
			var a = []byte("bbbb")

			got := LengthOfLongestSubstring(a)
			want := 1

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
			//ShouldNotBeNil(err)
		})

		Convey("求最长的无重复字符串长度,无重复", func() {
			var a = []byte("abcdefg")

			got := LengthOfLongestSubstring(a)
			want := 7

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
			//ShouldNotBeNil(err)
		})

	})

}
