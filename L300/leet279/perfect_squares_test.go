package leet279

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPerfectSquares(t *testing.T) {

	Convey("TestPerfectSquares", t, func() {

		Convey("1", func() {

			got := PerfectSquares(12)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := PerfectSquares(13)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}

func TestPerfectSquares1(t *testing.T) {

	Convey("TestPerfectSquares1", t, func() {

		Convey("1", func() {

			got := PerfectSquares1(12)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			got := PerfectSquares1(13)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
