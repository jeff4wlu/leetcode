package leet72

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEditDistance(t *testing.T) {

	Convey("TestEditDistance", t, func() {
		Convey("用例1", func() {

			got := EditDistance("horse", "ros")
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := EditDistance("intention", "execution")
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
