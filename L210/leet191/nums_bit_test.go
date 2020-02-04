package leet191

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumsBit(t *testing.T) {

	Convey("TestNumsBit", t, func() {
		Convey("用例1", func() {

			got := NumsBit(11)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
