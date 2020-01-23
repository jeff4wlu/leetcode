package leet151

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseWord(t *testing.T) {

	Convey("TestReverseWord", t, func() {
		Convey("用例1", func() {

			got := ReverseWord("the sky is blue")
			want := "blue is sky the"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("前后有空格，中间", func() {

			got := ReverseWord("  the    sky is blue  ")
			want := "blue is sky the"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("前有空格", func() {

			got := ReverseWord("  the   sky is blue")
			want := "blue is sky the"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("后有空格", func() {

			got := ReverseWord("the sky   is blue  ")
			want := "blue is sky the"

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
