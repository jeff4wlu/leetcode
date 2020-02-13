package leet256

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaintHouse(t *testing.T) {

	Convey("TestPaintHouse", t, func() {

		Convey("2", func() {

			got := PaintHouse([][]int{{14, 2, 11}, {11, 14, 5}, {14, 3, 10}})
			want := 10

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
