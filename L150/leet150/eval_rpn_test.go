package leet150

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEvalRPN(t *testing.T) {

	Convey("TestEvalRPN", t, func() {
		Convey("用例1", func() {

			in := []string{
				"2", "1", "+", "3", "*",
			}

			got := EvalRPN(in)
			want := 9

			if got != want {
				t.Errorf("failed")
			}

		})
		Convey("用例2", func() {

			in := []string{
				"4", "13", "5", "/", "+",
			}

			got := EvalRPN(in)
			want := 6

			if got != want {
				t.Errorf("failed")
			}

		})
		Convey("用例3", func() {

			in := []string{
				"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
			}

			got := EvalRPN(in)
			want := 22

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
