package leet204

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountPrimes(t *testing.T) {

	Convey("TestCountPrimes", t, func() {
		Convey("用例1", func() {

			got := CountPrimes(10)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := CountPrimes(100)
			want := 25

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
