package leet233

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumDigitOne(t *testing.T) {

	Convey("TestNumDigitOne", t, func() {
		Convey("1", func() {

			got := NumDigitOne(13)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
