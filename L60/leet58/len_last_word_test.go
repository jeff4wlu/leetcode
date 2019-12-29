package leet58

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLenLastWord(t *testing.T) {

	Convey("TestLenLastWord", t, func() {
		Convey("用例1", func() {
			s := "hello world"

			got := LenLastWord(s)
			want := 5

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("最后是空格", func() {
			s := "hello world   "

			got := LenLastWord(s)
			want := 5

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("全空格", func() {
			s := "   "

			got := LenLastWord(s)
			want := 0

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
