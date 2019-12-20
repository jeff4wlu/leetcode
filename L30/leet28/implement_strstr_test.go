package leet28

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImplementStrstr(t *testing.T) {

	Convey("TestImplementStrstr", t, func() {
		Convey("成功匹配", func() {

			haystack := "hello"
			needle := "ll"

			got := ImplementStrstr(haystack, needle)
			want := 2

			if got != want {
				t.Errorf("got is %d,want is %d", got, want)
			}

		})

		Convey("完全相同的情况", func() {

			haystack := "bba"
			needle := "bba"

			got := ImplementStrstr(haystack, needle)
			want := 0

			if got != want {
				t.Errorf("got is %d,want is %d", got, want)
			}

		})

		Convey("提前退出，省时间。不需要跑到最后才知道没有匹配的", func() {

			haystack := "aaaaabb"
			needle := "bba"

			got := ImplementStrstr(haystack, needle)
			want := -1

			if got != want {
				t.Errorf("got is %d,want is %d", got, want)
			}

		})

		Convey("一开始就匹配", func() {

			haystack := "aaaaabb"
			needle := "aa"

			got := ImplementStrstr(haystack, needle)
			want := 0

			if got != want {
				t.Errorf("got is %d,want is %d", got, want)
			}

		})

	})

}
