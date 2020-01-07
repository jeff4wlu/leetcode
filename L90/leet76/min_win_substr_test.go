package leet76

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinWinSubstr(t *testing.T) {

	Convey("TestMinWinSubstr", t, func() {
		Convey("用例1", func() {

			got := MinWinSubstr("ADOBECODEBANC", "ABC")
			want := "BANC"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

		Convey("用例2", func() {

			got := MinWinSubstr("ADBOBECODEBANC", "ABBC")
			want := "ADBOBEC"

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}

		})

	})

}
