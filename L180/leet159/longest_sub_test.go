package leet159

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPro(t *testing.T) {

	Convey("TestLongestSub", t, func() {
		Convey("1", func() {
			got := LongestSub("eceba")
			want := 3

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("2", func() {
			got := LongestSub("ccaabbb")
			want := 5

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
