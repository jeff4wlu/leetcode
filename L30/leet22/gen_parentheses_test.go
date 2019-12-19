package leet22

import (
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

			if got == nil || !comp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}

func comp(a, b []string) bool {
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}

	}
	return true
}
