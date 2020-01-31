package leet166

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFractionDec(t *testing.T) {

	Convey("TestFractionDec", t, func() {
		Convey("用例1", func() {

			got := FractionDec(-1, 2)
			want := "-0.5"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("用例2", func() {

			got := FractionDec(123, 999)
			want := "0.(123)"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("用例3", func() {

			got := FractionDec(2, 1)
			want := "2"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		//本题目不需要混排的循环小数
		Convey("用例4", func() {

			got := FractionDec(109, 900)
			want := "0.12(1)"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}
