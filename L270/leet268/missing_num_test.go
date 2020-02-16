package leet268

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMissingNum(t *testing.T) {

	Convey("TestMissingNum", t, func() {

		Convey("1", func() {

			got := MissingNum([]int{0, 1, 3})
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})
	})

}

func TestMissingNum1(t *testing.T) {

	Convey("TestMissingNum1", t, func() {

		Convey("1", func() {

			got := MissingNum1([]int{0, 1, 3})
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})
	})

}
