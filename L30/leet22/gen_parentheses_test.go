package leet22

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenParentheses(t *testing.T) {

	Convey("TestGenParentheses", t, func() {
		Convey("常规", func() {
			got := GenParentheses(3)
			want := []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}

			if got == nil || !infra.StringArrComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
