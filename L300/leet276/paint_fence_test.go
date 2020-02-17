package leet276

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaintFence(t *testing.T) {

	Convey("TestPaintFence", t, func() {

		Convey("1", func() {

			got := PaintFence(3, 2)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
