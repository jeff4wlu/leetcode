package leet258

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddDigits(t *testing.T) {

	Convey("TestAddDigits", t, func() {

		Convey("2", func() {

			got := AddDigits(38)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
