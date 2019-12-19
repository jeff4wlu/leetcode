package leet20

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidParentheses(t *testing.T) {

	Convey("TestValidParentheses", t, func() {
		Convey("常规1", func() {
			one := "()"
			got := ValidParentheses(one)
			want := true

			if got != want {
				t.Errorf("failed")

			}

		})

		Convey("常规2", func() {
			one := "()[]{}"
			got := ValidParentheses(one)
			want := true

			if got != want {
				t.Errorf("failed")

			}

		})

		Convey("常规3", func() {
			one := "(]"
			got := ValidParentheses(one)
			want := false

			if got != want {
				t.Errorf("failed")

			}

		})

		Convey("常规4", func() {
			one := "([)]"
			got := ValidParentheses(one)
			want := false

			if got != want {
				t.Errorf("failed")

			}

		})

		Convey("常规5", func() {
			one := "{[]}"
			got := ValidParentheses(one)
			want := true

			if got != want {
				t.Errorf("failed")

			}

		})

	})

}
