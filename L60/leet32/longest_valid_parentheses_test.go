package leet32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestValidParenthese(t *testing.T) {

	Convey("TestLongestValidParenthese", t, func() {
		Convey("用例1", func() {
			s := ")()())"
			got := LongestValidParenthese(s)
			want := 4

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("用例2", func() {
			s := "(()"
			got := LongestValidParenthese(s)
			want := 2

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("出现两个有效串，求出最长的", func() {
			s := "())()())"
			got := LongestValidParenthese(s)
			want := 4

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
