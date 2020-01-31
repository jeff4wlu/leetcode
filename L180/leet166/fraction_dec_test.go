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

			got := FractionDec(2, 3)
			want := "0.(6)"

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

	})

}
