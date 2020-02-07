package leet227

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicCal(t *testing.T) {

	Convey("TestBasicCal", t, func() {
		Convey("1", func() {

			got := BasicCal("3+2*2")
			want := 7

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := BasicCal(" 3/2 ")
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("3", func() {

			got := BasicCal(" 3+5 / 2 ")
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("4", func() {

			got := BasicCal(" 3+5 / 2 +2*3+1 ")
			want := 12

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
