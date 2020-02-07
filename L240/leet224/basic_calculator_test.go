package leet224

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicCal(t *testing.T) {

	Convey("TestBasicCal", t, func() {
		Convey("1", func() {

			got := BasicCal("1 + 1")
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := BasicCal(" 2-1 + 2 ")
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("3", func() {

			got := BasicCal("(1+(4+5+2)-3)+(6+8)")
			want := 23

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
